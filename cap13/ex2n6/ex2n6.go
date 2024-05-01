// - Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
// - Passe um valor do tipo slice of int como argumento para a função;
// - Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
// - Passe um valor do tipo slice of int como argumento para a função.
// - Solução: https://play.golang.org/p/Tgv3wwxKV-
package main

import "fmt"

func main() {
	r1 := vari(1, 2, 3, 4, 5, 6, 7, 8, 9)
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r2 := sliceInts(s1)

	fmt.Println(r1)

	fmt.Println(r2)

}

func vari(nums ...int) int {
	total := 0

	for _, v := range nums {
		total += v
	}

	return total
}

func sliceInts(nums []int) int {
	total := 0

	for _, v := range nums {
		total += v
	}

	return total
}
