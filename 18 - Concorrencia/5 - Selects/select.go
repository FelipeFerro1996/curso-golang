package main

import (
	"fmt"
	"time"
)

func main() {

	//declarado dois canais e enquanto um recebe valor a cada meio segundo um recebe a cada dois segundos, dessa forma utilizando select (Parecido com switch, porem so usado com concorrencia) posso usar a concorrencia da melhor forma possivel

	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	// for {
	// 	mensagem := <-canal1
	// 	fmt.Println(mensagem)

	// 	mensagem2 := <-canal2
	// 	fmt.Println(mensagem2)
	// }

	for {
		select {
		case mensagem := <-canal1:
			fmt.Println(mensagem)
		case mensagem2 := <-canal2:
			fmt.Println(mensagem2)
		}
	}

}
