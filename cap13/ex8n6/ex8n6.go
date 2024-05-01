// - Crie uma função que retorna uma função.
// - Atribua a função retornada a uma variável.
// - Chame a função retornada.
// - Solução: https://play.golang.org/p/A74rufv6Rs
package main

import "fmt"

func main() {
	dobrar := multi(2)

	r1 := dobrar(10)

	fmt.Println(r1)
}

func multi(multiplicador int) func(num int) int {
	return func(num int) int {
		return multiplicador * num
	}
}
