package main

import "fmt"

func main() {
	var_numero_inteiro := 10
	var_texto := "valor do meu texto"
	var_logica := true

	fmt.Println(
		fmt.Sprintf("o valor da var_numero_inteiro é : %v \n", var_numero_inteiro),
		fmt.Sprintf("o valor da var_texto é : %v \n", var_texto),
		fmt.Sprintf("o valor da var_logica é : %v \n", var_logica),
	)
}


/*
Obs: a formatacao é usada com a funcao fmt.Sprintf() que retorna o valor formatado.

### Palavras Chave para usar nas Formatacoes:
Formatacao qualquer tipo de valor/any  = %v

Formatacoes_Especificas
- formatacao string  = %s
- formatacao numero inteiro  = %d
- formatacao numero dizendo quantas casas decimais  = %.2f // diz que tera 2 casas decimais


*/