package main

import "fmt"

type endereco struct {
	rua    string
	numero int16
}

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

func main() {
	fmt.Println("Structs")

	var u usuario
	u.nome = "Felipe"
	u.idade = 30

	fmt.Println(u)

	enderecoExemplo := endereco{rua: "Sergio faria", numero: 344}

	usuario2 := usuario{"Felipe Ferro", 21, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Felipe Ferro ferro"}
	fmt.Println(usuario3)

}
