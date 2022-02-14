package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Números inteiros
	// Inteiro com sinal. int8 int16 int32 int64 (o tamanho depende da arquitetura da máquina)
	fmt.Println("Literal inteiro é", reflect.TypeOf(3200))

	// Inteiro sem sinal, só positivos. uint8 uint16 uint32 uint64
	// byte é um apelido para uint8
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// O rune é a representação na tabela Unicode e um alias para int32
	// É possível declarar com aspas simples caracteres únicos.
	var c rune = 'a'
	c2 := 'a'
	fmt.Println("O valor de 'a' é", c, "e o tipo é", reflect.TypeOf(c))
	fmt.Println("O valor de 'a' é", c2, "e o tipo é", reflect.TypeOf(c2))

	// ======================================================================================

	// Números reais (float32 e float64)
	// Por default ele usa o 64
	var d float32 = 49.99
	var e = 49.99
	fmt.Println("O tipo é", reflect.TypeOf(d))
	fmt.Println("O tipo é", reflect.TypeOf(e))

	// ======================================================================================

	// Boolean
	var f bool = true
	fmt.Println("Booleano", f)
	fmt.Println("Booleano negado", !f)

	// ======================================================================================

	// String
	// Aspas duplas
	var g string = "Stringzona"
	fmt.Println("String q pode ser concatenada usando o + " + g)
	// Para saber o tamanho real da strin, deve ser convertida para rune (q é a representação da string
	// pela tabela unicode). Caso o contrário, pode dar problema com acentos e Ç.
	fmt.Println("É possível saber o tamanho real convertendo para rune e usando o len", len([]rune(g)))

	// Com crase
	var h string = `Olha
	a quebra`
	fmt.Println("String", h)
	fmt.Println("É possível saber o tamanho usando o len", len(h))

	// ======================================================================================

	// Valores iniciais dos tipos, valores zero
	var i int
	var j float64
	var k bool
	var l string
	var m *int

	fmt.Println("Valores zero")
	fmt.Printf("int %v \nfloat %v \nbool %v \nstring %q \nponteiro %v\n", i, j, k, l, m)

	// ======================================================================================

	// Interface
	// Variáveis do tipo interface, são variáveis genéricas que aceitam valores de qualquer
	// tipo.
	fmt.Println("Tipo interface")

	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	coisa = "Stringzona"
	fmt.Println(coisa)

	// Tbm é possível criar um tipo
	type dinamico interface{}

	var coisa2 dinamico = "Wee"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

}
