// - Callback: passe uma função como argumento a outra função.
// - Solução: https://play.golang.org/p/2epLD_Yyap

package main

import "fmt"

func main() {
	exec(fala)
}

func exec(i func()) {
	i()
}

func fala() {
	fmt.Println("Hello World")
}
