package main

import "fmt"

// Estruturas lembram classes, é um agrupamento de dados
type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// É possível criar structs aninhadas
type compra struct {
	codigo   int
	produtos []produto
}

// Método são funções que possuem um receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

// Quando a struct possui uma struct, vc tem acesso aos "filhos" dela
// inclusive os métodoss
func (c compra) valorTotal() float64 {
	total := 0.0

	for _, prod := range c.produtos {
		total += prod.precoComDesconto()
	}

	return total
}

func main() {
	// Criando uma variável do tipo produto e iniciando ela
	// Dessa forma é possível acessar os métodos relacionados a essa estrutura
	var produto1 produto
	produto1 = produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}

	fmt.Println(produto1.precoComDesconto())

	// Declarando de forma simplificada uma variável do tipo produto
	produto2 := produto{"Notebook", 2000.50, 0.10}
	fmt.Println(produto2.precoComDesconto())

	// Declarando uma variável com struct aninhado
	compra1 := compra{
		codigo: 1,
		produtos: []produto{
			produto1,
			produto2,
		},
	}

	fmt.Println(compra1)
	fmt.Println(compra1.valorTotal())
}
