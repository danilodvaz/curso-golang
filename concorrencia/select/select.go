package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func maisRapido(url1, url2, url3 string) string {
	c1 := titulo(url1)
	c2 := titulo(url2)
	c3 := titulo(url3)

	// O Select é uma estrutura própria para utilizar em canais. É basicamente um
	// switch para ser utilizado em canais
	// Ele entrará na condição do primeiro canal que tiver uma resposa
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Estoutou o tempo de 1 segundo"
	}
}

func main() {
	rapido := maisRapido(
		"https://www.youtube.com/",
		"https://www.9gag.com",
		"https://www.google.com",
	)

	fmt.Println(rapido)
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
