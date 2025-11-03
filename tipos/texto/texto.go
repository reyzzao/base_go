package main

import "fmt"

func main() {
	meuTexto := "Sou meu texto"

	primeiroCaractere := meuTexto[0] // retorna a posicao na tabela Asci do caractere
	fmt.Println(primeiroCaractere)

	fmt.Println(meuTexto) // textos somente dentro de aspas duplas em golang.
	// var char = 'Oi' // erro: nao compila, Se adicionar dentro de aspas simples nem compila antigamente dava o numero da tabela Asci do caractere.

	// fmt.Println(char)
}
