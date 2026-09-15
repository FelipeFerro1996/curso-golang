package main

import "fmt"

func primeiraFuncao() {
	fmt.Println("Primeira função")
}

func segundaFuncao() {
	fmt.Println("Segunda função")
}

func retornaAlunoAprovado(n1, n2 int) bool {
	defer fmt.Println("Resultado calculado aguarde: o resultado já vaii ser informado")
	fmt.Println("Entrando na função para calculo da média")
	media := (n1 + n2) / 2

	if media >= 7 {
		return true
	}

	return false
}

func main() {

	defer primeiraFuncao()
	segundaFuncao()

	fmt.Println(retornaAlunoAprovado(7, 7))

}
