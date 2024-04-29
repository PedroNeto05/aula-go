// - Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
// - Demonstre os valores do map utilizando range.
// - Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
// - Solução: https://play.golang.org/p/GLK11Q1_x8y

package main

import "fmt"

type pessoa struct {
	name      string
	lastName  string
	fIceCream []string
}

func main() {

	mymap := make(map[string]pessoa)

	mymap["Neto"] = pessoa{
		name:      "Pedro",
		lastName:  "Neto",
		fIceCream: []string{"Chocolate", "Morango", "Menta com Chocolate"},
	}

	mymap["Nascimento"] = pessoa{
		name:      "Pedro",
		lastName:  "Nascimento",
		fIceCream: []string{"Laranja", "Uva", "Menta"},
	}

	for k, v := range mymap {
		fmt.Println(k, v)
		for k1, v1 := range v.fIceCream {
			fmt.Println(k1, v1)
		}
	}
}
