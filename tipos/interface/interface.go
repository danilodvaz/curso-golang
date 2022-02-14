package main

import "fmt"

type imprimivel interface {
	toString() string
	setNome(string)
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// Interfaces são implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func (p *pessoa) setNome(nome string) {
	p.nome = nome
}

func (p *produto) setNome(nome string) {
	p.nome = nome
}

// Como o método imprimir recebe uma variável do tipo interface,
// é possível passar qualquer struct q tenha os mesmos métodos definidos
// na interface, desde que respeite o padrão do método
func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func alteraNome(novoNome imprimivel, nome string) {
	novoNome.setNome(nome)
}

func main() {
	// A variável criado do tipo interface, tbm pode receber qualquer struct que
	// implemente aquela interface

	// Como a interface possui um método que recebe o endereçamento, é utilizado o
	// & para passar o endereço
	var algo imprimivel = &pessoa{"Roberto", "Silva"}
	fmt.Println(algo.toString())
	alteraNome(algo, "joaquim")
	imprimir(algo)

	algo = &produto{"Bermuda", 80.9}
	imprimir(algo)

	produto := &produto{"Calça Jeans", 179.00}
	fmt.Println(*produto)
	alteraNome(produto, "Bermudão")
	imprimir(produto)
}
