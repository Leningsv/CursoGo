package main

import (
	"fmt"
	"io/ioutil"
)

/*
	Esta libreria carga los datos en memoria,
	la lectura se la hace a nivel de bites, ya que el lector
	trabaja con bytes
*/
func main() {
	data, err := ioutil.ReadFile("./hola.txt")

	if err != nil {
		fmt.Println("Hubo un error")
	}
	// bites
	fmt.Println(data)
	// bites to string
	fmt.Println(string(data))
}
