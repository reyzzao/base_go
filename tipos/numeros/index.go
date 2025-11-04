package main

import "fmt"

// inteiros
var num_inteiro_any int = -170     // somente int deixa ser negativo
var num_inteiro_positivo uint = 33 // uint : não deixa ser negativo

// reais
var num_real float64 = 33.334

func main() {
	fmt.Println(num_inteiro_any)
	fmt.Println(num_real)
	fmt.Println(num_inteiro_positivo)
}
