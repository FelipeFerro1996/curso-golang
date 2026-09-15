package main

import (
	"fmt"
	"introducao-a-testes/enderecos"
)

func main() {
	endereco := enderecos.RetornatipoEndereco("Rodovia dos imigrantes")
	fmt.Println(endereco)
}
