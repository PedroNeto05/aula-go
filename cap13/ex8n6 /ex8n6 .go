// - Crie uma função que retorna uma função.
// - Atribua a função retornada a uma variável.
// - Chame a função retornada.
// - Solução: https://play.golang.org/p/A74rufv6Rs
package main

import "fmt"

func main() {
	dobra := multi(2)

	r := dobra(10)

	fmt.Println(r)
}

func multi(multiplicador int) func(int) int {
	return func(num int) int {
		return multiplicador * num
	}
}
