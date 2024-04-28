// - Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.
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
	delete(x, "haya")
	fmt.Println()
	for index, value := range x {
		fmt.Println(index, value)
	}
}
