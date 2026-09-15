package main

import "fmt"

var n int

func init() {
	n = 10
	fmt.Println("Executando função init")
}

func main() {

	fmt.Println("Executando função main")
	fmt.Println(n)

}
