// - Crie um tipo struct "pessoa" que contenha os campos:
//     - nome
//     - sobrenome
//     - idade
// - Crie um método para "pessoa" que demonstre o nome completo e a idade;
// - Crie um valor de tipo "pessoa";
// - Utilize o método criado para demonstrar esse valor.
// - Solução: https://play.golang.org/p/GBZcnu0Ajp

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) data() string {
	return fmt.Sprintln("Meu nome é", p.nome, p.sobrenome, "e tenho", p.idade, "anos")
}

func main() {
	pedro := pessoa{
		nome:      "Pedro",
		sobrenome: "Nascimento",
		idade:     19,
	}

	r1 := pedro.data()

	fmt.Println(r1)
}
