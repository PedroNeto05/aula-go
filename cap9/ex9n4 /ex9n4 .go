// - Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.
// - Solução: https://play.golang.org/p/3fcvHlt8Lm
package main

import "fmt"

func main() {
	x := map[string][]string{
		"Pedro": {"teste", "asdasdasd", "www"},
		"Neto":  {"teste", "asdasdasd", "www"},
		"Teste": {"teste", "asdasdasd", "www"},
		"Foo":   {"teste", "asdasdasd", "www"},
	}

	x["haya"] = []string{"teste", "asdasdasd", "www"}
	for index, value := range x {
		fmt.Println(index, value)
	}
}
