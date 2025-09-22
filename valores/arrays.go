package valores

import "fmt"

var arrayVazio = make([]string, 10)

func Arrays() {
	novoArray := append(arrayVazio, "Teste", "Teste2")
	fmt.Println(arrayVazio)
	fmt.Println(novoArray)
	fmt.Println(len(novoArray))
	fmt.Println(cap(novoArray))
}

/*
criar algum valor zerado: Usar: funcao make( tipo, tamanho, capacidade?)

- adicionarItem no Array: usar funcao append( arrayAlvo, items... )
- retorna tamanho do Array: usar funcao len( arrayAlvo )
- retorna capacidade do Array: usar funcao cap( arrayAlvo )
*/
