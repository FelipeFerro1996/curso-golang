package servidor

import (
	"crud-basico/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
