package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) teste() {
	fmt.Printf("O usuário %s foi salvo no banco \n", u.nome)
}

func (u usuario) maiorIdade() bool {
	return u.idade >= 18
}

// Ao declarar com ponteiro tudo que acontecer vai ser refletido fora
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {

	fmt.Println("Metodos")

	usuario1 := usuario{
		"Felipe",
		30,
	}
	fmt.Println(usuario1)
	usuario1.teste()

	usuario2 := usuario{
		"nayara",
		17,
	}
	fmt.Println(usuario2.maiorIdade(), usuario2)

	usuario1.fazerAniversario()

	fmt.Println(usuario1)
}
