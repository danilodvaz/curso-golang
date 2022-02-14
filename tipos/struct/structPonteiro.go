package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

// Como no Go são passado cópias dos valores como parâmetros, para realizar
// a alteração dos valores é ncessário fazer o apontamento para o endereço
// da memória
func (p *pessoa) setNome(nome, sobrenome string) {
	p.nome = nome
	p.sobrenome = sobrenome
}

func main() {
	pessoa1 := pessoa{"Tommy", "Shelby"}
	fmt.Println(pessoa1.getNomeCompleto())

	pessoa1.setNome("Thomas", "**** Shelby")
	fmt.Println(pessoa1.getNomeCompleto())
}
