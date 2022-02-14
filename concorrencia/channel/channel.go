package main

import (
	"fmt"
	"time"
)

// O canal é um ponto de sincronismo.
// Nele é possível enviar e receber informações
// Funciona como um fila
func main() {
	// Criando um canal de inteiros
	ch := make(chan int, 1)

	ch <- 1           // Enviando dados para canal (escrita)
	fmt.Println(<-ch) // Recebendo dados do canal(leitura)

	// Exemplo com routine
	ch2 := make(chan int)

	go exemploCanal(2, ch2)
	fmt.Println("Ta rodando a routine")

	a, b := <-ch2, <-ch2
	fmt.Println(a, b)

	c := <-ch2
	fmt.Println(c)
}

func exemploCanal(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(2 * time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}
