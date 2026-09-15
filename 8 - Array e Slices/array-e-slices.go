package main

import "fmt"

func main() {

	var array [5]int
	fmt.Println(array)

	array2 := [3]string{
		"posição 1",
		"Posição 2",
		"Posição 3",
	}
	fmt.Println(array2)

	array3 := [5]int{1, 2, 3, 4, 5}
	array3[3] = 7
	fmt.Println(array3)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(slice)

	slice = append(slice, 155)
	fmt.Println(slice)

	slice2 := array2[1:2]
	fmt.Println(slice2)

	array2[1] = "Teste de alteração"
	fmt.Println(slice2)

	fmt.Println("-----------------------")

	//Ao estourar o tamanho definido para um slice ele automaticamente gera uma nova posição e dobra a capacidade do mesmo para se adaptar aos novos parametros
	slice4 := make([]float32, 10, 11)
	slice4 = append(slice4, 10)
	slice4 = append(slice4, 15)

	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice5 := make([]float32, 5) //Ao não passar o parametro da capacidade do slice ele automaticamente pega o tamanho

	slice5 = append(slice5, 152)
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))
}
