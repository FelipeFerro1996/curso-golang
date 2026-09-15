package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j += 5 {
		fmt.Println("Incrementando J", j)
		time.Sleep(time.Second)
	}

	nomes := []string{"Felipe", "Nayara", "Neuza"}

	for index, value := range nomes {
		fmt.Println(index, value)
	}

	fmt.Println("--------------------")

	for index, value := range "Palavra" {
		fmt.Println(index, string(value))
	}
}
