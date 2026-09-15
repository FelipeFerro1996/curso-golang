package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = -100
	fmt.Println(numero)

	//unsygned
	var numero2 uint32 = 100000
	fmt.Println(numero2)

	//alias
	//int32 = RUNE
	var numero3 rune = 15222
	fmt.Println(numero3)

	//alias
	//uint8 = byte
	var numero4 byte = 152
	fmt.Println(numero4)

	//Numeros Reais

	var numerReal1 float32 = 150.00
	fmt.Println(numerReal1)

	var numerReal2 float64 = 150000000.05
	fmt.Println(numerReal2)

	numerReal3 := 150.00
	fmt.Println(numerReal3)

	//Strings

	var string1 = "teste"
	fmt.Println(string1)

	string2 := "teste 2"
	fmt.Println(string2)

	//o mais próximo que temos do char no go é onde o mesmo mostra o numero da tabela asc quando colocado o caractere dentro de aspas simples
	char := 'A'
	fmt.Println(char)

	//Ao não instanciar com um valor as variaveis todos os tipos possuem o seu '0' no caso do bool = "false"
	var booleano1 bool
	fmt.Println(booleano1)

	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)
}
