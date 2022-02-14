package main

import "fmt"

func main() {
	// Array
	// Array são estruturas com o mesmo tipo e com o tamanho fixo
	var notas [3]float64
	notas[0], notas[1], notas[2] = 2.3, 3.4, 5.4

	fila := [3]int{1, 2, 3}
	fmt.Println(fila)

	// Outra forma de declarar um array, onde o tamanho será fixo
	// e definido pela quantidade de itens que iniciaram a variável
	numeros := [...]int{1, 2, 3, 4, 5}
	fmt.Println(numeros)

	// ============================================================================

	// Slice
	// Slices são estruturas com tamanho variável.
	// Não são arrays. O slice referência um array interno.
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	// Criando slice com função make
	// Cria um slice apontando para um array interno de mesmo tamanho do slice
	slice2 := make([]int, 10)
	fmt.Println(slice2)

	// Cria um slice apontando para 10 elementos de um array iterno com 20 elementos
	slice3 := make([]int, 10, 20)
	fmt.Println(slice3, len(slice3), cap(slice3))

	// Adicionar elementos ao slice
	slice3 = append(slice3, 1, 2, 3)
	fmt.Println(slice3, len(slice3), cap(slice3))

	// Copiando um slice para o outro
	// Caso a origem seja menor que o destino, copia só as primeiras posições
	copy(slice, slice2)
	fmt.Println(slice, len(slice), cap(slice))

	// ============================================================================

	// Maps
	// Possui chave e valor, como se fosse um array associativo
	// Precisa ser inicializado
	aprovados := make(map[int]string)

	aprovados[123] = "Maria"
	aprovados[456] = "Pedro"
	fmt.Println(aprovados)

	// Outra forma de criar um map
	reprovados := map[string]float64{
		"João":  11.65,
		"Paulo": 45.32,
	}
	fmt.Println(reprovados)

	// Map aninhado
	todoMundo := map[string]map[string]float64{
		"M": {
			"Maria": 10,
		},
		"P": {
			"Pedro": 8,
			"Paulo": 7,
		},
		"J": {
			"João": 6,
		},
	}
	fmt.Println(todoMundo)

	// Deletando valores de um map
	delete(aprovados, 123)
	fmt.Println(aprovados)
}
