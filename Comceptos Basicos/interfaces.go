package main

import (
	"fmt"
)

type User interface {
	Permisos() int // 1 - 5
	Nombre() string
}

type Admin struct {
	nombre string
}

func (this Admin) Permisos() int {
	return 5
}
func (this Admin) Nombre() string {
	return this.nombre
}

type Editor struct {
	nombre string
}

func (this Editor) Permisos() int {
	return 3
}

func (this Editor) Nombre() string {
	return this.nombre
}

func auth(user User) string {
	permisos := user.Permisos()
	if permisos > 4 {
		return user.Nombre() + " Tiene permisos de administrador"
	} else if permisos == 3 {
		return user.Nombre() + " Tiene permisos de Editor"
	}
	return ""
}

/*
	Aunque no se implemnte un ainterfast basta con tener los
	mismos metodos para que se considere que una estructura implementa una interfas.
	Ademas una interfaz define un tipo de dato
*/
func main() {
	admin := Admin{"lenin"}
	editor := Editor{"fulano"}
	usuarios := []User{admin, editor}
	fmt.Println(auth(admin))
	fmt.Println(auth(editor))

	for _, usuario := range usuarios {
		fmt.Println("Ciclo", auth(usuario))
	}
}
