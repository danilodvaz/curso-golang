package main

import "fmt"

type carro struct {
	nome string
	ano  int
}

// Nesse caso, a struct ferrari vai agregar toda a struct carro, podendo
// acessar o que estiver em carro
// O resultado ao imprimir é o mesmo que declarando um dado do tipo carro,
// mas ao acessar não é cessário utilizar o nome do dado
type ferrari struct {
	carro
	modelo string
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.ano = 2010
	f.modelo = "Conversível"

	fmt.Println(f)
}
