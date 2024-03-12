package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// go test -v
// go test -cover
// go test -coverprofile cobertura.txt
// go tool cover -func=cobertura.txt
// go tool cover -html=cobertura.txt
// HTML output written to /tmp/cover12128398/coverage.html#file0
// go test -coverprofile cobertura.txt && go tool cover -html=cobertura.txt
// Só consegui abrir assim: file://wsl$/Ubuntu-20.04/tmp/cover1867538161/coverage.html

func TestTiposdeEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua das Rosas", "Rua"},
		{"Avenida das Rosas", "Avenida"},
		{"Estrada das Rosas", "Estrada"},
		{"Rodovia das Rosas", "Rodovia"},
		{"Praça das Rosas", "Tipo inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retorno := TipoDeEndereco(cenario.enderecoInserido)
		if retorno != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado. Esperava '%s', obteve '%s'", cenario.retornoEsperado, retorno)
		}
	}

}

func TestQualquer(t *testing.T) {
	if 10 < 5 {
		t.Error("Teste que sempre falha")
	}
}
