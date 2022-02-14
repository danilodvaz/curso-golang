package main

import "fmt"

func main() {
	a := 1
	b := 1.9999
	c := false
	d := "weee"

	// Quebra linha
	fmt.Println("")

	// Não quebra linah
	fmt.Print("")

	// Retorna uma string
	fmt.Sprint("")

	// Formata e retorna uma string
	fmt.Sprint("")

	// Printf formata o valor do print
	// Cada letra significa um tipo diferente
	fmt.Printf("\n %d %f %.1f %t %s", a, b, b, c, d)
	// Existe a letra v, que server para todos os tipos
	fmt.Printf("\n %v %v %v %v", a, b, c, d)
}
