package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "lalala"
		variavel4 string = "lalalala"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5", "variavel 6"

	fmt.Println(variavel5, variavel6)

	const constante string = "Contante 1"

	fmt.Println(constante)

	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5, variavel6)
}
