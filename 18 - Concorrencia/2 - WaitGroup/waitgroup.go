package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//Wait group ajuda a executar funções em paralelo, e o programa fique esperando as funções encerrarem para encerrar o mesmo
	var waitGroup sync.WaitGroup

	// Adicionado ao waitgroup que precisa esperar duas goroutines ser encerradas para encerrar o programa
	waitGroup.Add(2)

	go func() {
		escrever("Olá Mundo")
		waitGroup.Done() //-1
	}()

	go func() {
		escrever("Segunda Wait Group")
		waitGroup.Done() //-1
	}()

	//Indica ao Go que precisa aguardar as goroutines serem encerradas
	waitGroup.Wait()

}

func escrever(texto string) {

	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}

}
