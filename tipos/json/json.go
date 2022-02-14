package main

import (
	"encoding/json"
	"fmt"
)

// Da mesma forma que as funções, componentes dos struct iniciados com letra
// maiúscula é público e com letra minúscula é privado.
// O nome na frente de 'json' é nome que será utilizado como as chaves ao
// exportar para json
type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preço"`
	Tags  []string `json:"tags"`
}

func main() {
	p1 := produto{
		ID:    1,
		Nome:  "Notebook",
		Preco: 5_986.79,
		Tags:  []string{"Promoção", "Eletrônico"},
	}

	// Convertendo para json e retorna um slice de bytes
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// Json para struct
	stringJson := `{"id":2,"nome":"Monitor","preço":899.65,"tags":["Eletrônico"]}`
	// Converte uma slice de bytes para um struct
	var p2 produto
	json.Unmarshal([]byte(stringJson), &p2)
	fmt.Println(p2)
}
