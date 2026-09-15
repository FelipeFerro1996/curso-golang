package main

import "fmt"

func somarvariosNumeros(numeros ...int) int {
	total := 0

	for _, value := range numeros {
		total += value
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, value := range numeros {
		fmt.Println(texto, value)
	}
}

func main() {

	fmt.Println("Funções Variáticas")

	soma := somarvariosNumeros(1, 2, 3, 4, 5, 6, 7, 8, 155)
	fmt.Println(soma)

	escrever("Olá mundo: ", 10, 11, 1556, 1256)

}
