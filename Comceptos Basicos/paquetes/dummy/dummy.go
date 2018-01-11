/*
	Las libreiras no se puede jeecutar de manera independiente
*/
package dummy

var NombreDistinto string

func init() {
	NombreDistinto = "Hola"
}

func HolaMundo() string {
	return NombreDistinto
}

func holaMundo2() string {
	return "Soy privada no me pueden llamar"
}
