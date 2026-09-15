package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("teste")
	fmt.Println(resultado)

	//è possivel usar dois retornos na mesma função
	resultado1, resultado2 := calculosMatematicos(10, 15)

	fmt.Println(resultado1, resultado2)

	//é possivel ignorar um dos retornos passando um _ para ignorar o retorno especifico respeitando a ordem
	_, resultado3 := calculosMatematicos(10, 15)
	fmt.Println(resultado3)
}
