// - Crie e use um struct anônimo.
// - Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
// - Solução: https://play.golang.org/p/iTGLyH0Ijc & https://play.golang.org/p/h247Kid5adG

package main

import "fmt"

func main() {
	aStruct := struct {
		names map[string]string
		ages  []int
	}{
		names: map[string]string{
			"ola": "ola",
		},
		ages: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
	}

	fmt.Println(aStruct)
}
