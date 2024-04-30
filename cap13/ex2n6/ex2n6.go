// - Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
// - Passe um valor do tipo slice of int como argumento para a função;
// - Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
// - Passe um valor do tipo slice of int como argumento para a função.
// - Solução: https://play.golang.org/p/Tgv3wwxKV-

package main

import "fmt"

func main() {
	sliceInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r1 := slices(sliceInt)

	fmt.Println(r1)
	r2 := vari(1, 2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Println(r2)

}

func vari(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func slices(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}
