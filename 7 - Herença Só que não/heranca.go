package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint16
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança ó que não")

	p1 := pessoa{nome: "Felipe", sobrenome: "Ferro", idade: 30, altura: 182}

	e1 := estudante{pessoa: p1, curso: "Análise e desenvolvimento de sistema", faculdade: "Fatec garça"}

	fmt.Println(e1)
	fmt.Println(e1.nome, e1.altura)

}
