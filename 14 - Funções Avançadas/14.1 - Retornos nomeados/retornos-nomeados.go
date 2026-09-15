package main

import "fmt"

func calcumlosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	fmt.Println("Retornos Nomeados")

	soma, subtracao := calcumlosMatematicos(10, 20)
	fmt.Println(soma, subtracao)
}
