package main

import (
	"fmt"
	"time"
)

func main() {

	// Canais são usados para receber valores e travarem até que esse valores forem consumidos
	canal := make(chan string)

	go escrever("Felipe", canal)

	// for {
	// 	mensagem, aberto := <-canal
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Encerrado")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		// fmt.Println(texto)

		canal <- texto

		time.Sleep(time.Second)
	}

	close(canal)
}
