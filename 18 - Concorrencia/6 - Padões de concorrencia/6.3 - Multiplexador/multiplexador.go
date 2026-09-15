package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	canal1 := escrever("Olá Mundo")
	canal2 := escrever("Programando e go")

	canalsaida := multiplexador(canal1, canal2)

	for i := 0; i < 10; i++ {
		fmt.Println(<-canalsaida)
	}

}

func multiplexador(canalEntrada1, canalEntrada2 <-chan string) <-chan string {

	canal := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalEntrada1:
				canal <- mensagem
			case mensagem := <-canalEntrada2:
				canal <- mensagem
			}
		}
	}()

	return canal

}

func escrever(texto string) <-chan string {

	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("valor Recebido %s", texto)
			time.Sleep(time.Microsecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal

}
