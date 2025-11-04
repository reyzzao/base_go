package main

import "fmt"

/*
Fluxo Composicao de Entidade
1. Passo: 1, Estrutura_Pai, via: struct, exemplo: Animal
2. Passo: 2, Estruturas_Filhas, via: struct, exemplo: Gato & Cachorro
3. Passo: 3, Polimorfismo_Contrato_de_Acoes_Herdados_do_Pai, via: interface, exemplo: IAnimal { EmitirSom() }
4. Passo: 4, Polimorfismo_Funcoes_Construtoras_Fa, via: funcao, exemplo: FazerBarulho_FnConstrutora , sera usada quando quiser dar comportamento as Instancias_Filhas
5. Passo: 5, Polimorfismo: Metodo_dos_TiposFilhos_que_Implementam_o_Contrato_de_Acoes_Herdados_do_Pai
6. Passo: 6, Usar_Instanciar e popular os TiposFilhos
7. Passo: 7, usar Metodos_Construtores_da_Interface,  argumentando a instancia do TiposFilhos que desejar.
*/

// -------------------------------------------------------------

// Passo: 1, TipoPai , Perguntese: TODOS SERÃO UM
type Animal struct {
	Nome  string
	Idade uint
}

// Passo: 2, Estrutura_TiposFilhos : perguntese: SÃO UM ou Serão gerados pelo Pai, exemplo: Cachorro, Gato que SÃO um Animal
type Cachorro struct {
	Animal        // composicao: insere o pai, perguntese: este Objeto é UM?
	Cor    string // o que tera de diferente dos irmãos
}

// Passo: 2, Estrutura_TiposFilhos : perguntese: SÃO UM ou Serão gerados pelo Pai, exemplo: Cachorro, Gato que SÃO um Animal
type Gato struct {
	Animal        // composicao: insere o pai, perguntese: este Objeto é UM?
	Cor    string // o que tera de diferente dos irmãos
}

/*
 Passo: 3, Polimorfismo_Contrato_de_Acoes_Herdados_do_Pai :
 Sera um contrato de Contrato_de_Acoes_Herdados_do_Pai que todos Filhos terão, o nome da acao sera o mmesmo para todos, só sera executada de forma diferente por cada.
Palavra_Chave: `interface` : estrutura configuradora de acoes para estruturas
passo: 1 do polimorfismo (definir a interface)
*/

type IAnimal interface {
	// TODOS filhos poderam FAZER - Estes são Contrato_de_Acoes_Herdados_do_Pai
	EmitirSom()
}

/*
Passo: 4 , Metodos_Construtores_da_Interface
via: FuncaoConstrutora que recebe o Pai e devolve : Contrato_de_Acoes_Herdados_do_Pai
*/
func FazerBarulho_Metodos_Construtores_da_Interface(a IAnimal) {
	a.EmitirSom() // diz que o Pai deixou de heranca Contrato_de_Acoes_Herdados_do_Pai -então essa será AcaoDosFilhos_HerdadaDoPai
}

/*
Passo: 5, Metodo_dos_TiposFilhos_que_Implementam_o_Contrato_de_Acoes_Herdados_do_Pai
*/
func (g Gato) EmitirSom() {
	fmt.Println("Miauuuuu")
}

func (g Cachorro) EmitirSom() {
	fmt.Println("Au..Au")
}

func main() {

	/* Passo: 6, Usar_Instanciar e popular os TiposFilhos */
	gato1 := Gato{
		Animal: Animal{Nome: "gato1", Idade: 10},
		Cor:    "Marron",
	}

	cachorro1 := Cachorro{
		Animal: Animal{Nome: "cachorro1", Idade: 8},
		Cor:    "Dourado",
	}

	/*
		Passo: 7, Usar os Metodos_Construtores_da_Interface
		- passando as Instancias_Filhas
		- Uso: Esta será a funcao que poderão usar nas instancias, não desencadeando mas sim chamando a função e passando a instancia.

	*/
	FazerBarulho_Metodos_Construtores_da_Interface(gato1)     // usando o metodo vindo de interface na intancia, passando a infancia
	FazerBarulho_Metodos_Construtores_da_Interface(cachorro1) // usando o metodo vindo de interface na intancia, passando a infancia

}
