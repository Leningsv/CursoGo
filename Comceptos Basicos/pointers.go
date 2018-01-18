package main

import (
	"fmt"
)

/*
	1. Un puntero es una direccion de memoria
	2. Se tiene la direccion en la que esta el valor
	3. X,Y => #dir_mem => 5
	*T => *int, *string, *Struct
*/
func main() {
	var x, y *int
	entero := 5
	x = &entero
	y = &entero
	*x = 6
	fmt.Println(x, entero)
	fmt.Println(y, entero)
}
