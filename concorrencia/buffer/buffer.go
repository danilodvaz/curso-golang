package main

import "fmt"

func rotina(ch chan int) {
	ch <- 1
	fmt.Println("Executando 1")
	ch <- 2
	fmt.Println("Executando 2")
	ch <- 3
	fmt.Println("Executando 3")
	ch <- 4
	fmt.Println("Executando 4")
	ch <- 5
	fmt.Println("Executando 5")
	ch <- 6
	fmt.Println("Executando 6")
}

func main() {
	// Criando um canal com 3 buffers
	// O buffer define um tamanho para o canal
	// Quando o buffer enche, o código fica esperando liberar espaço para adicionar
	// mais um valor no canal
	ch := make(chan int, 3)
	go rotina(ch)

	fmt.Println(<-ch)

	// É possível utilizar estruturas de repetição em um canal
	ch2 := make(chan int)
	go repeticaoCanal(ch2)

	// Quando for fazer algo desse tipo, lembrar de colocar um close no canal para
	// evitar deadlock
	for numero := range ch2 {
		fmt.Println(numero)
	}

	fmt.Println("terminou")
}

func repeticaoCanal(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
	}

	close(c)
}
