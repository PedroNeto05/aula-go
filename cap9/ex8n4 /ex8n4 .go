// - Crie um map com key tipo string e value tipo []string.
//     - Key deve conter nomes no formato sobrenome_nome
//     - Value deve conter os hobbies favoritos da pessoa
// - Demonstre todos esses valores e seus índices.
// - Solução: https://play.golang.org/p/nD3TW8VQmH

package main

import "fmt"

func main() {
	x := map[string][]string{
		"Pedro": {"teste", "asdasdasd", "www"},
		"Neto":  {"teste", "asdasdasd", "www"},
		"Teste": {"teste", "asdasdasd", "www"},
		"Foo":   {"teste", "asdasdasd", "www"},
	}

	for index, value := range x {
		fmt.Println(index, value)
	}
}
