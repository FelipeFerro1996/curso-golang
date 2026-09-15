package main

import "fmt"

// ao usar worker varias vezes lendo os canais, criamos uma função que recebe os numero e outra que processa dessa forma podemos usar varias vezes em paralelo, aumentando o desempenho da função
func main() {

	tarefas, resultados := make(chan int, 45), make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}

	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

}

func worker(tarefas <-chan int, resultados chan<- int) {

	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}

}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
