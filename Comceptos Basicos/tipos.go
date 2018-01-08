package main

// import "fmt"
import (
	"fmt"
	"strconv"
)

func main() {
	edad := 22
	edadStr1 := "13"
	edadStr := strconv.Itoa(edad)
	// las funciones de go pueden retornar mas de un valor.
	// _ es un comodin que sirve para declarar una variable que se desechara y no sera usada
	edadNumero, _ := strconv.Atoi(edadStr1)
	fmt.Println("Mi edad es " + edadStr)
	fmt.Println(edadNumero + 100)
}
