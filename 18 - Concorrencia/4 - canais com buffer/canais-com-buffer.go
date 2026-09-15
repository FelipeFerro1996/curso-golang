package main

import "fmt"

func main() {

	//canais com buffer indica que o canal pode receber dois valores antes de travar a execução do canal
	canal := make(chan string, 2)

	canal <- "Olá Mundo"
	canal <- "Olá Mundo2"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
