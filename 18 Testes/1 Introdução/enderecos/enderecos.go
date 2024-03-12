package enderecos

import "strings"

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"Rua", "Avenida", "Estrada", "Rodovia"}
	primeiraPalavra := strings.Split(endereco, " ")[0]

	enderecoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			enderecoValido = true
		}
	}

	if enderecoValido {
		return primeiraPalavra
	}

	return "Tipo inválido"
}
