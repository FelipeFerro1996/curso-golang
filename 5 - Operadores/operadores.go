package main

import "fmt"

func main() {
	soma := 1 + 2
	subtracao := 2 - 1
	divisao := 10 / 3
	multiplicacao := 3 * 2

	fmt.Println(soma, divisao, subtracao, multiplicacao)

	fmt.Println("--------------")
	numero := 10

	numero--
	numero++
	numero += 12
	numero -= 10
	numero *= 3
	numero /= 2

	fmt.Println(numero)

	var texto string

	if numero > 5 {
		texto = "Número é maior do que 5"
	} else {
		texto = "Númer não é maior do que 5"
	}

	fmt.Println("---------------")
	fmt.Println(texto)

	fmt.Println("---------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro || falso)
	fmt.Println(verdadeiro && falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	fmt.Println("---------------")
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
}
