// - Crie um loop utilizando a sintaxe: for {}
// - Utilize-o para demonstrar os anos desde que você nasceu.
// - Solução: https://play.golang.org/p/dIbfdiuw0ms

package main

import "fmt"

func main() {
	i := 2004
	for {

		fmt.Println(i)
		i++
		if i > 2024 {
			break
		}
	}
}
