package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	Obtiene al archivo como un objeto. Se tiene que
	abrir y cerrar el archivo.
	defer: Se ejectua luego que termina una funcion sea el resultado
		bueno u error
	panic: Imprime un error con la finalidad de saber en que linea se da el error.
		busca el origen del error la cascada hasta el origen. Paniquea el programa
		no se ejecutara nada a partir de la generacion del error
	recover: Detiene el paniqueo, y continua la ejecucion normal, a partir del
		lugar de su invocacion. recover contiene el error que paniqueo el script.
*/
func main() {
	executeReadFile()
	fmt.Println("Nunca me ejecutare")
}

func executeReadFile() {
	ejecucion := readFile()
	fmt.Println(ejecucion)
}

func readFile() bool {
	archivo, error := os.Open("./hola1.txt")
	// Super util. Se ejecuta siempre al final
	defer func() {
		archivo.Close()
		fmt.Println("Defer", "Me jecute al final")
		r := recover()
		fmt.Println(r)
	}()
	if error != nil {
		panic(error)
	}
	scanner := bufio.NewScanner(archivo)
	var i int
	for scanner.Scan() {
		i++
		linea := scanner.Text()
		fmt.Println(i, linea)
	}
	if true {
		return true
	}
	fmt.Println("Nunca me ejecuto")
	return true
}
