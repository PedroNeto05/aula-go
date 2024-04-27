// - Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
// - Solução: https://play.golang.org/p/4-iTPZwfEz

package main

import "fmt"

func main() {
	esporteFavorito := "asdasdasd"

	switch esporteFavorito {
	case "Futebol":
		fmt.Println(esporteFavorito)
	case "Vólei":
		fmt.Println(esporteFavorito)
	case "Basquete":
		fmt.Println(esporteFavorito)
	case "Handball":
		fmt.Println(esporteFavorito)
	default:
		fmt.Println("Nenhum")
	}

}
