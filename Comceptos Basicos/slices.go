package main

import (
	"fmt"
)

//El valor nulo en GO es un valor el elemmento inicializado pero vacio
/*
	La estructura de un slice contiene:
	- Longitud
	- Puntero al Slice(Arreglo)
	- Capacidad
*/
func main() {
	slice := []int{1, 2, 3, 4}
	var slice1 []int
	arreglo := [3]int{1, 2, 3}
	slice2 := arreglo[:]
	fmt.Println(slice, "Capacidad:", cap(slice))
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(slice2)
}
