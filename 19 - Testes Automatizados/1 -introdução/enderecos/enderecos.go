package enderecos

import (
	"strings"
)

// RetornatipoEndereco Verifica se o endereço tem um tipo válido e o retorna
func RetornatipoEndereco(endereco string) string {

	tiposValidos := []string{"rua", "avenida", "rodovia"}

	enderecoMinisculo := strings.ToLower(endereco)

	primeiraPalavraDoEndereco := strings.Split(enderecoMinisculo, " ")[0]

	enderecotesteTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecotesteTipoValido = true
		}
	}

	if enderecotesteTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"

}
