package main

import "fmt"

type Conta struct {
	Titular string
	Numero  string
	Saldo   float64
}

// metodo: acoes com os campos que tem na estrutura , Perguntese: o que a estrutura pode FAZER
func (c *Conta) Credito(valor float64) {
	c.Saldo += valor
}

/*
Funcao_Construtora
- sempre retorna um ponteiro da estrutura desejada
- pede em params somente valores não automatizados
- devolve valores default para a estrutura
*/
func NovaConta(titular string) *Conta {
	return &Conta{
		Titular: titular,
		Numero:  "12345",
		Saldo:   0.0,
	}
}

func main() {
	// composicao literal
	// conta1 := Conta{Titular: "Reinaldo", Numero: "1", Saldo: 0.00} // usando literal marretado
	conta2 := NovaConta("Titular 2") // usando Funcao_Construtora ,argumentando parametros pedidos
	conta2.Credito(20)
	// visualizacoes
	// fmt.Println("Situação atual da conta1 >> ", conta1)
	fmt.Println("Situação atual da conta2 >> ", conta2)
}
