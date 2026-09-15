package main

import "fmt"

func inverterNumero(numero int) int {
	return numero * -1
}

func inverterNumeroComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {

	numero := 20

	numeroInvertido := inverterNumero(numero)
	fmt.Println(numeroInvertido, numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterNumeroComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
