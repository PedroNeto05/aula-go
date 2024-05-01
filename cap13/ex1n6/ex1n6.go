// - Exercício:
//     - Crie uma função que retorne um int
//     - Crie outra função que retorne um int e uma string
//     - Chame as duas funções
//     - Demonstre seus resultados
// - Solução: https://play.golang.org/p/rxJM5fgI-9

package main

import "fmt"

func main() {
	r1 := returnInt()
	r2, r3 := returnIntAndString()

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3) 
}

func returnInt() int {
	return 2 * 2
}

func returnIntAndString() (int, string) {
	return 2 * 2, "Ola"
}
