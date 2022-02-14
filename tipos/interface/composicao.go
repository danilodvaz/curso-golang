package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

// A interface esportivoLuxuoso é uma composição das interfaces esportivo e luxuoso
// Ela tem acesso a todos os métodos das outras interfaces
type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// pode adicionar mais métodos
}

type bmw struct{}

func (b bmw) ligarTurbo() {
	fmt.Println("Ligando turbo")
}

func (b bmw) fazerBaliza() {
	fmt.Println("Fazendo baliza")
}

func main() {
	var carro esportivoLuxuoso = bmw{}
	carro.fazerBaliza()
	carro.ligarTurbo()
}
