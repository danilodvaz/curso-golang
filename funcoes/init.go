package main

import "fmt"

// Dentro de cada arquivo do Go, pode existir a função init,
// que vai iniciar alguma coisa
func init() {
	fmt.Println("Inicializando")
}

func main() {
	fmt.Println("Main")
}
