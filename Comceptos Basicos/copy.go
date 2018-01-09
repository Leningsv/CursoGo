package main

import (
	"fmt"
)

/*
	Copia el minimo de elementos en ambos arreglos
*/
func main() {
	fuente := []int{1, 2, 3, 4}
	destino := make([]int, len(fuente), cap(fuente)*2)
	fmt.Println(fuente)
	fmt.Println(destino)
	copy(destino, fuente)
	fmt.Println(fuente)
	fmt.Println(destino)
}
