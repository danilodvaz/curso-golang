package main

import "math"

// Quando iniciado com letra maiúscula, significa que é público dentro do pacote
// Quando iniciado com letra minúscula, significa que é privado dentro do pacote

type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)

	return
}

func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)

	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
