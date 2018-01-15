package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	edad := 22
	var edad1 int
	bander := true
	precio := 13.35
	// %v valor original del dato
	// %t imprime boleanos
	// %b imprime boleanos
	fmt.Printf("Hola Mundo %d", edad)
	fmt.Printf("Hola Mundo %t", bander)
	fmt.Printf("Hola Mundo %b", precio)

	fmt.Println("\nColoca tu edad: ")
	// Toma un dato desde consola y lo asigna a una variable
	fmt.Scanf("%dc\n", &edad1)
	fmt.Printf("Mi edad es: %d", edad1)

	// Toma un dato desde afuera
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nIngresa tu nombre: ")
	nombre, err := reader.ReadString('0')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hola " + nombre)
	}

}
