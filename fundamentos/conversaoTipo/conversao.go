package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// Conversão para string

	// Pega a string da tabela asc
	fmt.Println("Teste " + string(97))

	// Int para string
	fmt.Println("Teste " + strconv.Itoa(97))

	// String para inteiro
	num, _ := strconv.Atoi("97")
	fmt.Println(reflect.TypeOf(num))
}
