package main

import "fmt"

func main() {
	fmt.Println(testandoDefer(10))
}

// O defer permite colocar alguma coisa para ser executada antes de qualquer
// return do método
func testandoDefer(numero int) int {
	defer fmt.Println("Executa antes do return do método")

	if numero > 10 {
		fmt.Println("Número maior que 10")
		return numero
	}

	fmt.Println("Número menor ou igual a 10")
	return numero
}
