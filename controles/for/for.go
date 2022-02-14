package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Println(i)
	}

	// for {
	// 	fmt.Println("Infinito")
	// 	time.Sleep(time.Second)
	// }

	numeros := [...]int{1, 2, 3, 4, 5}
	for index, valor := range numeros {
		fmt.Println("Index", index, "Valor", valor)
	}
}
