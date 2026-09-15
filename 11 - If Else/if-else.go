package main

import "fmt"

func main() {

	fmt.Println("Operadores de Controle")

	numero := 10

	if numero > 10 {
		fmt.Println("O Número é maior do que 10")
	} else {
		fmt.Println("O Número é menor ou igual do que 10")
	}

	if outronumero := numero; outronumero > 0 {
		fmt.Println("Outro número maior do que 0")
	} else if outronumero > -10 {
		fmt.Println("O número está entre 0 e -10")
	} else {
		fmt.Println("Número menor do que -10")
	}

}
