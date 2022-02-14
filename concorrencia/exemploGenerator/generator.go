package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// <-chan - A função irá retorna um canal somente leitura
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

func main() {
	t1 := titulo("https://www.udemy.com/", "https://www.youtube.com/")
	t2 := titulo("https://www.google.com", "https://www.9gag.com")

	// Os primeiros sites que forem retornando irão preencher o canal, quer será
	// lido enquanto é preenchido
	fmt.Println("Primeiro:", <-t1)
	fmt.Println("Segundo:", <-t1)
	fmt.Println("Terceiro:", <-t2)
	fmt.Println("Quarto:", <-t2)
}
