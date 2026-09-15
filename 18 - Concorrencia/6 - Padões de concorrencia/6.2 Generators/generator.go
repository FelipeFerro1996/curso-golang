package main

import (
	"fmt"
	"time"
)

func main() {

	// generators são funções que inicializam e criam um canal, retirando da função main essa responsabilidade
	canal := escrever("Olá Mundo")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}

}

func escrever(texto string) <-chan string {

	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Second)
		}
	}()

	return canal

}
