// - Crie um programa que demonstre o funcionamento da declaração if.
// - Solução: https://play.golang.org/p/OGPgTJZbiy

package main

import "fmt"

func main() {
	x := 10

	if x%2 == 0 {
		fmt.Println(x)
	}
}
