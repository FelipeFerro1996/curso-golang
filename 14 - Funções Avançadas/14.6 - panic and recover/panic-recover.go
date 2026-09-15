package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func alunoAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2
	fmt.Println("media = ", media)
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("O programa entrou em panico")
}

func main() {

	fmt.Println("Panic and Recover")

	fmt.Println(alunoAprovado(10, 4))

}
