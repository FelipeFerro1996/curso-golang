package main

import "fmt"

func main() {

	retorn := func(texto string) string {
		return fmt.Sprintf("Valor recebido -> %s", texto)
	}("Olá Mundo")

	fmt.Println(retorn)
}
