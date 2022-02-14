package main

import "fmt"

func main() {
	imprimirResultado(8)
	imprimirResultado(5.9)
	imprimirResultado(3)

	ifComInit()
}

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else if nota >= 4 && nota <= 6 {
		fmt.Println("Repescagem com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

// Permite iniciar uma variável dentro do escopo do if apenas
// Mesmo esquema do for. Pode ser utilizado uma função para definir essa variável
func ifComInit() {
	if i := 2; i > 1 {
		fmt.Println("Entrou", i)
	}
}
