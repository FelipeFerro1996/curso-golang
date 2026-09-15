package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {

	cachorroEmJson := `{"nome":"Nina","raca":"Indefinida","idade":6}`
	fmt.Println(cachorroEmJson)

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJson), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	c2 := make(map[string]string)

	cachorro2EmJson := `{"nome":"Toby", "raca":"Poodle"}`

	if erro := json.Unmarshal([]byte(cachorro2EmJson), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EmJson, c2)

}
