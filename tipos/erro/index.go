package main

import (
	"errors"
	"fmt"
)

// configurar a funcionalidade com tratamento:
func dividirDoisNumeros(dividendo, divisor int) (int, error) {

	if divisor == 0 {
		return 0, errors.New("O divisor nao pode ser 0")
		// tratar: é verificar/testar e devolver o tipo erro com seu valor zero e o feedback na instancia do erro.
	}

	return dividendo / divisor, nil
	// devolucao: a operacao que dara o resultado favoravel, e o nulo para ser preenchido pelo tratamento da verificacao que colocara a instancia do erro passado acima.
}

// exemplo se caso nao tratasse
func dividir_SemTratar(dividendo, divisor int) int {
	return dividendo / divisor
}

// usar a funcionalidade tratada
func main() {
	// verificar o que devolve de acerto e erro na funcionalidade e usa-los.
	result, err := dividirDoisNumeros(10, 0)
	if err != nil {
		fmt.Println(result)
	}
	fmt.Println(err.Error())

	// usar a funcionalidade sem tratamento.
	// fmt.Println(dividir_SemTratar(10, 0)) // se dividir sem tratar da panic() e o programa é encerrado, o panic pode ser tratado com recorevy()
}
