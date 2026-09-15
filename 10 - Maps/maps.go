package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario1 := map[string]string{
		"nome":      "Felipe",
		"Sobrenome": "Ferro",
		"idade":     "30 anos",
	}

	fmt.Println(usuario1, usuario1["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "felipe",
			"segundo":  "Ferro",
		},
		"endereco": {
			"rua":    "ergio Faria",
			"bairro": "Santa Antonieta",
		},
	}

	fmt.Println(usuario2)
}
