package main

import "fmt"

// variaveis constantes
// pode ser declarado dentro ou fora de funcao
const PI = 3.14

/*
	 constantes possuem dois tipos:
		digitadas
		nao digitadas
*/
const A int = 300 // digitada
const B = 500     // nao digitada; tipo de dado e inferido

// declarando constantes em bloco
const (
	C int = 1
	D = 3.14
	E = "Hi!"
  )

func main() {
	/* variaveis declarando o tipo

	   A declaração de variáveis ​​e a atribuição de valores podem ser feitas separadamente

	   Pode ser usado dentro e fora das funções
	*/
	var variavelInteira int = 10

	// declarar multiplas variaveis
	var a1, b1, c1, d1 int = 1, 3, 5, 7
	e, f := 7, "World!"

	/* variavel com inferencia de tipo
	   quando declarado tem que passar valor

	   A declaração de variáveis ​​e a atribuição de valores não podem ser feitas separadamente
	   (devem ser feitas na mesma linha)

	   Só pode ser usado dentro de funções
	*/
	variavelString := "helloWorld"

	/* variaveis nao inicializadas
	   mesmo sem valor inicial declarado
	   ainda tera um valor dito por valores padrao de cada tipo
	*/
	var a2 string
	var b2 int
	var c2 bool

	// declarar variaveis em bloco
	var (
		var1 int
		var2 int    = 1
		var3 string = "hello"
	)

	fmt.Println("variavel tipada (declarando tipo): ", variavelInteira)
	fmt.Println("variavel nao tipada (inferencia de tipo): ", variavelString)

	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(d1)

	fmt.Println(a2)
	fmt.Println(b2)
	fmt.Println(c2)

	fmt.Println(e)
	fmt.Println(f)

	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
}
