package main

import "fmt"

func main() {
	i := 1

	// Criando uma variável que irá receber o endereço de memória de um inteiro
	// Como ponteiros são o endereço da memória, ele pode ser nulo
	var p *int = nil

	// Pegando o endereço de memória da variável
	p = &i

	// acessando o valor do endereço
	*p++
	i++

	fmt.Println(p, *p, i, &i)
}
