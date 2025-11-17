package main

import (
	"fmt"
	"time"
)

func main() {
	// enquanto uma condição for verdadeira vou repetir o codigo
	// o iteravel vem de fora do for e é modificado dentro do for

	i := 0
	for i < 10 {
		i++
		fmt.Println(fmt.Sprintf("Incrementando i : %v", i))
		time.Sleep(time.Second)
	}
}
