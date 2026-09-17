package servidor

import (
	"crud-basico/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	corporequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario

	if erro := json.Unmarshal(corporequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao popular struct usuario"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios(nome,email) values(?,?)")
	if erro != nil {
		fmt.Println(erro)
		w.Write([]byte("Erro criar o statement"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao salvar dados no banco "))
		return
	}

	usuarioInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao retornar ultimo usuario inserido"))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário incluído com sucesso. ID: %d", usuarioInserido)))

}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuários no banco de dados"))
		return
	}

	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {

		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao processar as linhas de usuarios"))
			return
		}

		usuarios = append(usuarios, usuario)

	}

	w.WriteHeader(http.StatusOK)

	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter usuarios para json"))
		return
	}

}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parametro id para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	defer db.Close()

	linha, erro := db.Query("select * from usuarios where id = ?", ID)

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuario"))
			return
		}
	}
	defer linha.Close()

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuario para json"))
		return
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parametro ID para inteiro"))
		return
	}

	corporequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao processar o corpo da requisição"))
		return
	}

	var usuario usuario

	if erro := json.Unmarshal(corporequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao transformar em json o body"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(&usuario.Nome, &usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao executar o update no usuario"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parametro id para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		fmt.Println(erro)
		w.Write([]byte("Erro ao realizar o statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar o usuaŕio"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
