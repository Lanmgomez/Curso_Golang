package main

import "fmt"

func main() {

	nome, idade, vagabundo := Informacoes("Islan", 26, true)

	fmt.Println(nome, idade, vagabundo)

	// podemos declarar uma variavel que recebe funcoes

	resultadoSoma := func() {
		valor1 := 2
		valor2 := 1
		total := valor1 + valor2
		fmt.Println("O resultado da soma eh:", total)
	}

	resultadoSoma()

	resultadoSubtracao := func(a int, b string) (int, string) {
		return a, b
	}

	a, b := resultadoSubtracao(26, "Islan")

	fmt.Println(a, b)
}

// variacoes de funçoes
// funç~oes em Go tem que ter o tipo de retorno apos os argumentos
// o retorno pode ser um ou mais valores

func exemplo1(a int, b int) int {
	return a + b
}

func exemplo2(a int, b int) (int, string) {
	return a + b, "resultado:"
}

func Informacoes(nome string, idade int, vagabundo bool) (string, int, bool) {
	return nome, idade, vagabundo
}
