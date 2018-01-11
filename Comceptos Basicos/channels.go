package main

import (
	"fmt"
)

/*
	Envia informacion entre go rutines, el chanel espera a que resiva la informacion
	del canal. Aunque el funcionamiento es asicronico cuando se interactua sobre un canal
	existe una espera para que resiva informacio n y seguidamente se realiza una nueva
	ejecucion
*/
func main() {
	channel := make(chan string)
	go func(channel chan string) {
		// Este for siempre esta esperando
		for {
			var name string
			fmt.Scanln(&name)
			channel <- name
		}
	}(channel)

	msg := <-channel
	fmt.Println("Estoy imprimiendo lo que resivi del canal: " + msg)

	msg = <-channel
	fmt.Println("Estoy imprimiendo lo que resivi del canal: " + msg)
}
