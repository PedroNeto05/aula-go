// - Exercício:
//     - Crie uma função que retorne um int
//     - Crie outra função que retorne um int e uma string
//     - Chame as duas funções
//     - Demonstre seus resultados
// - Solução: https://play.golang.org/p/rxJM5fgI-9

package main

import "fmt"

func main() {
	x := returnInt()
	y, z := returnIntAndString()

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func returnInt() int {
	return 2 * 2
}

func returnIntAndString() (int, string) {
	return 2 * 2, "Ola"
}
