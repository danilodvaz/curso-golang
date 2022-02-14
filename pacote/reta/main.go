package main

// Os controles e acessos são realizados apenas no fora do pacote.
// Tudo dentro de um mesmo pacote é compartilhado.
// Quando iniciado com letra maiúscula, significa que é público dentro do pacote
// Quando iniciado com letra minúscula, significa que é privado dentro do pacote

import "fmt"

func main() {
	// É possível acessar tudo no mesmo pacote
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}

	fmt.Println(catetos(p1, p2))
	fmt.Println(Distancia(p1, p2))
}
