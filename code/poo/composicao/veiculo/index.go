package main

import "fmt"

type Veiculo interface {
	buzina() string
}

type Carro struct {
	Nome string
	Ano  int
	// Dono Pessoa // obs: nao pode fazer circulo viciador de composicao, Um tipo ter Um e este Um ter o tipo tambem.
}

func (c Carro) buzina() string {
	return fmt.Sprintf("Ok buzineiiiii o meu Veiculo chamado >> %v", c.Nome)
}

type Moto struct {
	Nome string
	Ano  int
}

func (m Moto) buzina() string {
	return fmt.Sprintf("Ok buzineiiiii o meu Veiculo chamado >> %v", m.Nome)
}

type Pessoa struct {
	Nome   string
	Compra Veiculo
}

func CreatePessoa(p Pessoa) *Pessoa {
	return &p
}

func CreateCarro(v Carro) Carro {
	return v
}
func CreateMoto(v Moto) Moto {
	return v
}

var Carro1 = CreateCarro(Carro{Nome: "NomeCarro1", Ano: 2000})
var Moto1 = CreateMoto(Moto{Nome: "NomeMoto1", Ano: 2000})

var Pessoa1 = CreatePessoa(Pessoa{Nome: "Nome_Reinaldo", Compra: Carro1})
var Pessoa2 = CreatePessoa(Pessoa{Nome: "Nome_Leo", Compra: Moto1})

func main() {

	// fmt.Println(Pessoa1)
	// fmt.Println(Pessoa1.Veiculo.buzina())

	// fmt.Println(Pessoa2)
	// fmt.Println(Pessoa2.Veiculo.buzina())

	fmt.Println(Pessoa2)
	fmt.Println(Pessoa2.Compra.buzina())
}
