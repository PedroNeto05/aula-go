// - Crie um loop utilizando a sintaxe: for condition {}
// - Utilize-o para demonstrar os anos desde que você nasceu.
// - Solução: https://play.golang.org/p/qnFjiDJzLor

package main

import "fmt"

func main() {
	i := 2004
	for i <= 2024 {
		fmt.Println(i)
		i++
	}
}
