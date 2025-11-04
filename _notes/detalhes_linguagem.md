
### Declaracoes

Podemos declarar varias variaveis/parametros alinhados e inserir seu tipo uma unica vez depois do nome da variavel 
exemplo :
``` 
func divisao ( dividendo, divisor int ) {}
// declarei varios params e por ultimo add o tipo de todos.
```

### Tipos

#### Texto
- palavra_chave_do_tipo: `string`

```
fmt.Println("Ola textos") 
// textos somente dentro de aspas duplas em golang.
// var char = 'Oi' // erro: nao compila, Se adicionar dentro de aspas simples nem compila antigamente dava o numero da tabela Asci do caractere.
```

### Funcao
- palavra_chave_do_tipo: ``func``

```

	// uma funcao em go sempre vem o resultado favoravel e o erro entao >>
	// tratar é verificar/testar e devolver o tipo erro com seu valor zero e o feedback na instancia do erro.
	// mas em go sempre temos que verificar se a funcao nao contem erro e se conter fazer algo com ele, se nao quiser testa-la use o anderline para nao fazer nada com o erro
	// no uso da funcao eu poderia somente executa-la mas vamos fazer o test e usar o favoravel e o que ela traz de erro,

```

### Erro
- palavra_chave_do_tipo: `err`
- objeto_para_uso_das_props_deste_tipo: `errors.`
```
	// nulo / nil	: é o valor default do tipo erro, é quando o espaço na memoria/variavel nao contem anda nenhuma informação.


```


### Ainda_Nao_Modular
- uma pasta só pode ter um arquivo com funcao main
- se o package for main e a funcao de saida tambem main() pode executar o arquivo.
