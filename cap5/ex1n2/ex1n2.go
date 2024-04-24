// - Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
// - Solução: https://play.golang.org/p/X7qm3aWSa6

package main

import "fmt"

func main() {
	x := 1
	y := 2
	z := 3

	fmt.Printf("%d\t %b\t %#X\n", x, x, x)
	fmt.Printf("%d\t %b\t %#X\n", y, y, y)
	fmt.Printf("%d\t %b\t %#X\n", z, z, z)
}
