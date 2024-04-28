// - Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
// - "Nome", "Sobrenome", "Hobby favorito"
// - Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
// - Solução: https://play.golang.org/p/Gh81-d5tMi

package main

import "fmt"

func main() {
	x := [][]string{
		{"Pedro", "Neto", "Programar"},
		{"Pedro", "Neto", "Jogar"},
		{"Pedro", "Neto", "Comer"},
	}

	for _, value := range x {
		for _, value1 := range value {
			fmt.Println(value1)
		}
		fmt.Println()
	}
}
