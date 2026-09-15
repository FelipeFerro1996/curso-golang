package main

import "fmt"

func diasDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "Número Inválido"
	}
}

func diadaSemana2(numero int) string {

	var diadaSemana string

	switch {
	case numero == 1:
		diadaSemana = "Domingo"
		fallthrough //Ao executar a instrução automaticamente passa para o outro case sem validar o mesmo
	case numero == 2:
		diadaSemana = "Segunda"
	case numero == 3:
		diadaSemana = "Terça"
	case numero == 4:
		diadaSemana = "Quarta"
	case numero == 5:
		diadaSemana = "Quinta"
	case numero == 6:
		diadaSemana = "Sexta"
	case numero == 7:
		diadaSemana = "Sabado"
	default:
		diadaSemana = "Número Inválido"
	}

	return diadaSemana
}

func main() {
	fmt.Println("Switch")

	dia := diasDaSemana(10)
	fmt.Println(dia)

	dia2 := diadaSemana2(1)
	fmt.Println(dia2)
}
