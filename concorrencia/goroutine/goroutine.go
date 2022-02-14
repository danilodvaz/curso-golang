package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// Execução sincrona
	fale("Maria", "Pq vc não fala comigo?", 3)
	fale("João", "Só posso falar depois de vc!", 1)

	// Ao adicionar a palavra reservada 'go', ele vai criar uma execução assíncrona
	// daquele código. Ele vai mandar o código executar e já pular para a próxima
	// linha.
	// É importante lembrar que tudo será finalizado, ao finalizar a execução da
	// função main
	go fale("Jiji", "Wee", 10)
	fale("Juju", "Abacate", 5)
}
