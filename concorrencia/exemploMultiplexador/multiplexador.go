package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Função que recebe dois canais e transfere os dados de um canal origem para
// um canal destino
// Recebe como parâmetro um canal somente leitura origem e um canal destino
func encaminhar(origem <-chan string, destino chan string) {
	// Acessa o canal origem e alimenta o destino.
	// Quando não houver nada no canal origem, o sistema ficará esperando.
	for {
		destino <- <-origem
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)

	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)

	return c
}

func main() {
	c := juntar(
		titulo("https://www.udemy.com/", "https://www.youtube.com/"),
		titulo("https://www.google.com", "https://www.9gag.com"),
	)

	fmt.Println(<-c, " | ", <-c)
	fmt.Println(<-c, " | ", <-c)
}

func titulo(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		// Criando uma função anônima e logo no final, já executando ela com os ()
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}
