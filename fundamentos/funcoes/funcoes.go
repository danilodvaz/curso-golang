package main

import "fmt"

func somar(a int, b int) int {
	return a + b
}

func imprimir(valor int) {
	fmt.Println(valor)
}

func retornoDefinido() (resultado string) {
	resultado = "Danilo"
	return
}

func main() {
	resultado := somar(3, 4)
	imprimir(resultado)

	fmt.Println(retornoDefinido())
}
