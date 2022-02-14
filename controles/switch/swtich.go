package main

import "fmt"

func main() {
	nota := 8
	estado := "reprovado"

	// Switch padrão
	switch nota {
	case 10:
		fallthrough //Faz executar o prórximo caso
	case 9:
		estado = "A"
	case 8, 7: //Mais de uma opção
		estado = "B"
	default:
		estado = "Inválido"
	}

	fmt.Println("Situação aluno", estado)

	// Switch true
	// Procura o primeiro resultado verdadeiro dentro dos cases
	switch {
	case 1 > 2:
		fmt.Println("Não entra")
	case 1 > 1:
		fmt.Println("Não entra")
	case 1 == 1:
		fmt.Println("Entra")
	default:
		fmt.Println("Não entra")
	}

	// Exemplo de switch para identificar o tipo da entrada
	var i interface{}
	var tipo string

	switch i.(type) {
	case int:
		tipo = "inteiro"
	case float32, float64:
		tipo = "real"
	case string:
		tipo = "string"
	case func():
		tipo = "função"
	default:
		tipo = "indefinido"
	}

	fmt.Println(tipo)
}
