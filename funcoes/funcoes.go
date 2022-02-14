package main

import "fmt"

func main() {

}

// Função sem parâmetro e sem retorno
func f1() {
	fmt.Println("weee")
}

// Função recebendo parâmetro
func f2(p1 string, p2 string) {
	fmt.Println(p1, p2)
}

// Função apenas com retorno
func f3() string {
	return "Wee"
}

// Função que recebe e retorna
// É igual (p1 string, p2 string)
func f4(p1, p2 string) string {
	return p1 + p2
}

// Função que retorna mais de um resultado
func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

// Retorno nomeado
// Tem o retorno limpo
func f6(p1, p2 int) (segundo, primeiro int) {
	segundo = p1
	primeiro = p2

	return
}

// Função anônima adicionada em uma variável
// Pode ser executada chamando a variável soma(2, 2)
var soma = func(a, b int) int {
	return a + b
}

// Utilizar uma função como parâmetro
// Função que recebe uma função como parâmetro e mais dois inteiros
// A variável de entrada 'funcao' é do tipo função, recendo dois inteiros e retornando um inteiro
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func multiplicacao(a, b int) int {
	return a * b
}

func funcaoParam() int {
	// É possível passar a função 'multiplicao' como parâmetro para a função exec
	return exec(multiplicacao, 1, 2)
}

// Função que retorna uma outra função
func f7() func() {
	var funcao = func() {
		fmt.Println("weee")
	}

	return funcao
}

// Funções variáticas (Tipo javascript com o spread)
// Funções com parâmetro variável
// Pode ser chamada da seguinte forma: media(2, 3, 6, 8)
// Tbm pode ser chamada passando um slice: media(slice...)
func media(numeros ...int) int {
	// A variável deve ser tratada como um array
	total := 0

	for _, num := range numeros {
		total += num
	}

	return total / len(numeros)
}

// Função recursiva
// Chamada da função: fatorial(3)
func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

// Função recebendo endereço de memória (ponteiro)
// Ele vai alterar diretamente no endereço da memória, atualizando a variável
// que foi passada
// Chamada da função: f9(&numero)
func f9(numero *int) {
	*numero++
}
