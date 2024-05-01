// - Atribua uma função a uma variável.
// - Chame essa função.
// - Solução: https://play.golang.org/p/RMHLL3N5Ww

package main

import "fmt"

func main() {
	x := func() {
		fmt.Println("OLA")
	}

	x()
}
