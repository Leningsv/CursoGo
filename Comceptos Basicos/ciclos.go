package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hola mundo")
	// Ciclo for
	for i := 0; i < 10; i++ {
		fmt.Println(i, "Hello World")
	}
	// Ciclo While
	j := 0
	for j < 10 {
		fmt.Println(j, "While")
		j++
	}

	// Ciclo While
	j = 0
	for {
		if j == 2 {
			j++
			continue
		}

		fmt.Println(j, "While - Break")
		j++

		if j >= 10 {
			break
		}
	}
}
