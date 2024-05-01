// - Demonstre o funcionamento de um closure.
// - Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.
// - Solução: https://play.golang.org/p/sA7NHpkCCg
package main

import "fmt"

func main() {

	dobra := multi(2)
	triplica := multi(3)
	quadriplica := multi(4)
	quituplica := multi(5)

	num := 10

	r1 := dobra(num)
	r2 := triplica(num)
	r3 := quadriplica(num)
	r4 := quituplica(num)

	fmt.Println(r1, r2, r3, r4)

}

func multi(multiplicador int) func(int) int {
	return func(num int) int {
		return multiplicador * num
	}
}
