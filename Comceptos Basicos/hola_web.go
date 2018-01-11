package main

import (
	"fmt"
	"io"
	"net/http"
)

/*
	w: representa un string de datos de escritura.
	se puede escribir funcionalidad para diferenctes URLs
*/
func main() {
	http.HandleFunc("/holamundo", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hola Home")
	})
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hay una nueva peticion")
	io.WriteString(w, "hola mundo")
}
