package main

import "fmt"

// Tipo_Pai , Perguntese: TODOS SERÃO UM
type Animal struct {
	Nome  string
	Idade uint
}

// Estrutura_Filha : perguntese: SÃO UM ou Serão gerados pelo Pai, exemplo: Cachorro, Gato que SÃO um Animal
type Cachorro struct {
	Animal        // composicao: insere o pai, perguntese: este Objeto é UM?
	Cor    string // o que tera de diferente dos irmãos
}

// Estrutura_Filha : perguntese: SÃO UM ou Serão gerados pelo Pai, exemplo: Cachorro, Gato que SÃO um Animal
type Gato struct {
	Animal        // composicao: insere o pai, perguntese: este Objeto é UM?
	Cor    string // o que tera de diferente dos irmãos
}

// passo: 3 do polimorfismo: cada Estrutura_Filha terá que implementar/fazer metodo do seu jeito com o Contrato_de_Acoes_Herdados_do_Pai

func (g Gato) EmitirSom() {
	fmt.Println("Miauuuuu")
}

func (g Cachorro) EmitirSom() {
	fmt.Println("Au..Au")
}

/*
 Polimorfismo :
 Sera um contrato de Contrato_de_Acoes_Herdados_do_Pai que todos Filhos terão, o nome da acao sera o mmesmo para todos, só sera executada de forma diferente por cada.
Palavra_Chave: `interface` : estrutura configuradora de acoes para estruturas
passo: 1 do polimorfismo (definir a interface)
*/

type IAnimal interface {
	// TODOS filhos poderam FAZER - Estes são Contrato_de_Acoes_Herdados_do_Pai
	EmitirSom()
}

/*
passo: 2 do polimorfismo : (implementar deixar a disposicao dos filhos para eles fazerem metodos com a funcao que recebera o Contrato_de_Acoes_Herdados_do_Pai)
AcaoDosFilhos_HerdadaDoPai - Configurando a Contrato_de_Acoes_Herdados_do_Pai
- que cada filho tera que implementar do seu jeito
- Perguntese: O Nome do Acao da interface é o MESMO QUE?
a resposta sera o nome da funcao o que os filhos farao.
Neural: para o Pai o Nome é mais Formal para os filhos será mais moderno ,
exemplo: EmitirSom() É O MESMMO QUE FazerBarulho
Obs: Em Go é diferente o uso de metodo na instancia se o metodo vir de interface: essa funcao sea usada quando quiser dar comportamento as Instancias_Filhas, ao invés de desencadear da instancia no uso tem que chamar a funcao e passar a instancia.
*/
func FazerBarulho(a IAnimal) {
	a.EmitirSom() // diz que o Pai deixou de heranca Contrato_de_Acoes_Herdados_do_Pai -então essa será AcaoDosFilhos_HerdadaDoPai
}

func main() {

	/* Usando_Instancias_Filhos : */
	// instanciando preenchendo valores
	gato1 := Gato{
		Animal: Animal{Nome: "gato1", Idade: 10},
		Cor:    "Marron",
	}

	cachorro1 := Cachorro{
		Animal: Animal{Nome: "cachorro1", Idade: 8},
		Cor:    "Dourado",
	}

	// Usando Instancias_Filhas_Ja_Instanciadas
	FazerBarulho(gato1)     // usando o metodo vindo de interface na intancia, passando a infancia
	FazerBarulho(cachorro1) // usando o metodo vindo de interface na intancia, passando a infancia

}
