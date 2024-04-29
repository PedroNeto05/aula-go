// - Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
//     - Nome
//     - Sobrenome
//     - Sabores favoritos de sorvete
// - Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
// - Solução: https://play.golang.org/p/Pyp7vmTJfY

package main

import "fmt"

type pessoa struct {
	name      string
	lastName  string
	fIceCream []string
}

func main() {
	pessoa1 := pessoa{
		name:      "Pedro",
		lastName:  "Neto",
		fIceCream: []string{"Chocolate", "Morango", "Menta com Chocolate"},
	}
	pessoa2 := pessoa{
		name:      "Pedro",
		lastName:  "Nascimento",
		fIceCream: []string{"Laranja", "Uva", "Menta"},
	}
	fmt.Println(pessoa1)
	for _, v := range pessoa1.fIceCream {
		fmt.Println(v)
	}
	fmt.Println(pessoa2)
	for _, v := range pessoa2.fIceCream {
		fmt.Println(v)
	}
}
