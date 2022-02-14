package main

import (
	"log"
	"net/http"
)

func main() {
	// Inicia um fileserver apontando para a pasta public
	fs := http.FileServer(http.Dir("public"))

	// Direciona o / para o fileserver
	http.Handle("/", fs)

	log.Println("Executando...")
	// Sobe o servidor na porta 3000
	log.Fatal(http.ListenAndServe(":3000", nil))
}
