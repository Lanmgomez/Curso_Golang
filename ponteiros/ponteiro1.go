package ponteiros

import "fmt"

func Ponteiro1() {

	valor := 1
	var ponteiroValor *int = &valor

	*ponteiroValor = 2

	valor2 := &valor
	*valor2 = 3

	fmt.Println("o valor do Ponteiro1:", *ponteiroValor)
}
