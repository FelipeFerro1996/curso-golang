package main

import (
	"aluno-aprovado/app"
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Projeto para gerar aluno aprovado em andamento")

	var nota1, nota2 float64
	var nome string

	fmt.Println(nota1, nota2)

	fmt.Println("Informe o nome do aluno que vai ser calculado a nota")
	_, err := fmt.Scanln(&nome)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Informe a primeira nota")
	_, err1 := fmt.Scanln(&nota1)

	fmt.Println("Informe a segunda nota")
	_, err2 := fmt.Scanln(&nota2)

	if err1 != nil || err2 != nil {
		fmt.Println("Erro: Informe as notas corretamente")
		return
	}

	media, aprovado := app.CalcularNota(nota1, nota2)

	fmt.Printf("O aluno %s, teve média %0.2f, portanto está:\n", nome, media)

	time.Sleep(time.Second)

	if aprovado {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}

}
