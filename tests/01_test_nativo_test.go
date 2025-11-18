package index_test

import (
	p "github/test_pro"
	"testing"
)

func TestSoma(t *testing.T) {
	if p.Soma(2, 2) != 41 {
		t.Errorf("========== Esta funcao Soma deve retornar >> 4  ==========================")
	} // 4)
}

/*
Obs: o arquivo da suite de test : no seu nome deve ter o prefixo final `_test.go`
Rodar_todos_tests: go test ./...
rodar_Test_Especifico: go test ./<pasta do arquivo alvo de test>
criar_arquivo_de_coberturas_de_test: `go test -coverprofile=coverage.out ./...`
criar_html_preview_dos_tests_cobertos: `go tool cover -html=coverage.out -o coverage.html`

*/
