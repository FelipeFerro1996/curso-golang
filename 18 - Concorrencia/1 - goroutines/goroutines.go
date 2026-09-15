package main

import (
	"fmt"
	"time"
)

func main() {
	//Concorrencia != paralelismo
	go escrever("Olá Mundo!") //goroutine - a função é dispárada e não é preciso que a mesma termine para continuar o codigo
	escrever("Programando em go")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
