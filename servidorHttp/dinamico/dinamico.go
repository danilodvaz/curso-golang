package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// O ResponseWriter é responsável para escrever no response
// O Request é a requisição
func horaCerta(w http.ResponseWriter, r *http.Request) {
	// Formata a data. Cada valor sendo passado tem uma função, representando
	// o dia, mês, ano, hora....
	s := time.Now().Format("02/01/2006 03:04:05")

	// Escreve no response
	fmt.Fprintf(w, "<h1>Hora certa: %s</h1>", s)
}

func main() {
	// Direciona para a função informada
	http.HandleFunc("/horaCerta", horaCerta)

	log.Println("Excecutando....")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
