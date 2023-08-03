package main

import (
	"fmt"
	"std/app/ponteiros"
)

func main() {

	valor := 10
	var ponteiro *int = &valor

	// tras o endereço da memoria: 0xc0000ac010
	fmt.Println(ponteiro)

	// tras o valor armazenado no ponteiro
	fmt.Println(*ponteiro)

	Ponteiros()

	ponteiros.Ponteiro1()

}

// Exemplo 2

func Ponteiros() {

	// a variavel começa com um valor
	teste := "Islan"
	var ponteiroTeste *string = &teste

	// apos criaçao do ponteiro, podemos reescrever o valor dela
	*ponteiroTeste = "Teste"

	novoValor := &teste
	*novoValor = "Valor Novo"

	fmt.Println("O endereço de memoria eh:", ponteiroTeste)
	fmt.Println("O valor armazenado eh:", *ponteiroTeste)
}
