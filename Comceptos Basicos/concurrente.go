package main

import (
	"fmt"
	"strings"
	"time"
)

/*
	Para decirle que la funcion se ejecute en bacground,
	Una ejecucion separada, concurrente.
	go rutine: numero es muy alto. Abre un nuevo treat cunado las goritunes estan
	bloquedas.
	go en funciones anonimas
*/
func main() {
	go miNombreLento("Lenin")
	fmt.Println("Que aburrido")
	var wait string
	fmt.Scanln(&wait)
}

func miNombreLento(name string) {
	letras := strings.Split(name, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
