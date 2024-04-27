// - Crie um programa que utilize a declaração switch, sem switch statement especificado.
// - Solução: https://play.golang.org/p/TyIGp4Hi8B

package main

import "fmt"

func main() {
	x := "asdasdasd"

	switch {
	case x == "Pedro":
		fmt.Println(x)
	default:
		fmt.Println("Ninguem")
	}

}
