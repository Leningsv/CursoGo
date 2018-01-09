package main

import (
	"fmt"
)

func main() {
	var arreglo [10]int
	arreglo1 := [4]int{1, 2}
	var matriz [2][2]int

	fmt.Println(arreglo, "Tamaño", len(arreglo))
	fmt.Println(arreglo1)

	for i := 0; i < len(arreglo1); i++ {
		fmt.Println(arreglo1[i])
	}
	matriz[0][1] = 23
	fmt.Println(matriz)

}
