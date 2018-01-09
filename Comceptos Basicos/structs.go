package main

import (
	"fmt"
)

type User struct {
	edad             int
	nombre, apellido string
}

func (usuario User) nombreCompleto() string {
	return usuario.nombre + " " + usuario.apellido
}

func (usuario User) setName(name string) {
	usuario.nombre = name
}
func (usuario *User) setNameP(name string) {
	usuario.nombre = name

}

/*
	Define la estructura de un tipo de dato
	todas las funciones en go duplican el objeto y lo envian
*/
func main() {
	var lenin User
	usuario := User{
		nombre:   "Lenin",
		apellido: "Samaniego",
		edad:     29}
	usuario1 := User{0, "nombre", "apellido"}
	usuario2 := new(User) //Retorna un puntero
	(*usuario2).nombre = "puntero - nombre"
	fmt.Println(lenin)
	fmt.Println(usuario)
	fmt.Println(usuario1)
	fmt.Println(usuario2)
	(*usuario2).nombre = "puntero - cambio nombre"
	fmt.Println(usuario2)
	usuario.setName("Lenin1")
	fmt.Println(usuario.nombreCompleto())
	usuario.setNameP("Lenin1")
	fmt.Println(usuario.nombreCompleto())
}
