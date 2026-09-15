package enderecos

import (
	"testing"
)

type cenarioDeteste struct {
	enderecoEnviado string
	retornoEsperado string
}

//ferramentas e comando para execução de testes
//go test -  executa o teste dentro da pasta aonde voce está
//go test ./... - roda todos os testes da pasta do projeto
//go test --cover - executa o teste e mostra a cobertura do teste, indicando se está totalmente testado ou não
//go test --coverprofile arquivo.txt -- realiza o teste e gera um arquivo com as informações do teste executado
//go tool cover --func=arquivo.txt - le o arquivo de teste gerado na pasta pois o mesmo não é interpretado sem o comando
//go tool cover --html=arquivo.txt - gera um arquivo html indicando quais linhas foram testadas e quais não foram da  minha função a ser testada

func TestTiposDeEnderecos(t *testing.T) {
	t.Parallel()

	enderecoParaTeste := "Avenida Paulista"

	tipoEnderecoEsperado := "Avenida"

	tipoEnderecoRetornado := RetornatipoEndereco(enderecoParaTeste)

	if tipoEnderecoEsperado != tipoEnderecoRetornado {
		t.Errorf(
			"O tipo de endereço retornado é diferente do esperado. Era esperado %s e foi retornado %s",
			tipoEnderecoEsperado,
			tipoEnderecoRetornado,
		)
	}

}

func TestTiposDeEnderecosStruct(t *testing.T) {

	t.Parallel()

	cenariostestes := []cenarioDeteste{
		{"Rua ABC", "Rua"},
		{"Avenida rebouças", "Avenida"},
		{"Rodovia Raposo tavares", "Rodovia"},
		// {"", "Tipo Inválido"},
		// {"Estrada Teste", "Tipo Inválido"},
		{"RUA ABC", "Rua"},
		{"AVENIDA RIO BRANCO", "Avenida"},
		{"RUA DAS PATATIVAS", "Rua"},
	}

	for _, cenario := range cenariostestes {
		retornorecebido := RetornatipoEndereco(cenario.enderecoEnviado)

		if retornorecebido != cenario.retornoEsperado {
			t.Errorf("Erro: tipo esperado %s, tipo recebido %s",
				cenario.retornoEsperado,
				retornorecebido,
			)
		}
	}

}
