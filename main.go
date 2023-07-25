package main

import (
	"fmt"
	"std/app/soma2"
)

// 01 - Tipagem "explicita"
var nome string
var idade int

func main() {
	nome = "Islan"
	idade = 26
	fmt.Println(`O nome eh:`, nome, `a idade eh:`, idade)

	// 02 - tipagem "direta"
	nome2 := "Lanm"
	idade2 := 26
	fmt.Println("O 2° nome eh:", nome2, "A idade 2 eh:", idade2)

	Soma1()

	resultado := soma2.Soma2(5, 5)
	fmt.Println("O valor da soma 2 eh:", resultado)
}

func Soma1() {
	num1 := 5
	num2 := 4
	total := num1 + num2
	fmt.Println("O valor da soma 1 eh:", total)
}
