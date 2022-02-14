package main

// É possível colocar um alias na importação do pacote
import (
	"fmt"
	m "math"
)

func main() {
	// Criando uma constante de forma padrão
	const PI float64 = 3.1415

	// Declarando mais de uma constante
	const (
		a = 1
		b = 2
	)

	// Criando uma variável de forma padrão
	var raio = 3.2 // Quando é definido uma valor, já é inferido o tipo

	// Declarando mais de uma variável
	var (
		c = 3
		d = 4
	)
	var e, f bool = true, false

	// Criando variável de forma reduzida (PREFERÊNCIA)
	area := PI * m.Pow(raio, 2)
	g, h, i := 2, false, "wee"

	fmt.Println("A área da circunferência é", area)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}
