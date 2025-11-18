package main

import "fmt"

// para iterar/percorrer ua estrutura, use a clausula/funcao range para percorrer a estrutura. e ter 2 vars para trabalhar
var nomes = []string{"Reinaldo", "Guga", "Leonardo", "Gabriel"}

var objetoPessoas = map[string]string{
	"key1": "Programador",
	"key2": "Sistemas",
}

func handleNomes() {
	// sintaxe: for var_posicao, var_valor será := um range em var_array { aqui faço o que quiser nas vars capturadas var_posicao, var_valor }
	for posicao, valor := range nomes {
		fmt.Println(posicao, valor)
	}
}

// no caso de objeto map, as variaveis disponiveis para iterar/percorrer são a chave e valor
func handlerObjetoPessoas() {
	for key, value := range objetoPessoas {
		fmt.Println(key, value)
	}
}

func main() {
	// handleNomes()
	handlerObjetoPessoas()

}
