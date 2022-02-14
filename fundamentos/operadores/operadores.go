package main

import "fmt"

func main() {
	a := true && true
	b := true != false
	c := true || false
	d := !true
	fmt.Println(a, b, c, d)

	// Operadores bitwise (bit a bit)
	e := 3
	f := 2

	fmt.Println("AND", e&f) // 11 & 10 = 10 (Apenas os bits em comum)
	fmt.Println("OR", e|f)  // 11 | 10 = 11 (O que tiver já vale)
	fmt.Println("XOR", e^f) // 11 ^ 10 = 01 (Ou exclusivo)
}
