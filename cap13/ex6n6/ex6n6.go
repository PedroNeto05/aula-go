// - Crie e utilize uma função anônima.
// - Solução: https://play.golang.org/p/Kgo6hVr5G5

package main

import "fmt"

func main() {
	func() {
		fmt.Println("OLA")
	}()

}
