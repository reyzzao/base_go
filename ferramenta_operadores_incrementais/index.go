package main

import "fmt"

func main() {
	// adicicao em 1 a cada execucao use na variavel ++ que é o mesmo que var = var + 1
	adicionar_em_1 := 0
	adicionar_em_1++
	adicionar_em_1++
	adicionar_em_1++
	fmt.Println(adicionar_em_1)

	adicionar_em_2 := 0
	adicionar_em_2 += 2
	adicionar_em_2 += 2
	adicionar_em_2 += 2
	fmt.Println(adicionar_em_2)

}
