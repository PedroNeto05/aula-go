// - Usando uma literal composta:
//     - Crie uma slice de tipo int
//     - Atribua 10 valores a ela
// - Utilize range para demonstrar todos estes valores.
// - E utilize format printing para demonstrar seu tipo.
// - Solução: https://play.golang.org/p/ST3TKusuOd

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, value := range slice {
		fmt.Println(index, value)
	}

	fmt.Printf("%T", slice)
}
