package main

import (
	"fmt"
	"time"
)

func main() {
	// sintaxe: o iteravel , a condicao e a mudanca do iteravel sao definidas na declaracao do for
	for i := 0; i < 10; i += 2 {
		fmt.Println(fmt.Sprintf("Incrementando i : %v", i))
		time.Sleep(time.Second)
	}
}
