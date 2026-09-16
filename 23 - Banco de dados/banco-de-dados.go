package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	stringConsexao := "go:go@/curso_go?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConsexao)
	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	if erro = db.Ping(); erro != nil {
		fmt.Println("Erro dentro do ping")
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}

	linhas.Close()

	fmt.Println(linhas)

}
