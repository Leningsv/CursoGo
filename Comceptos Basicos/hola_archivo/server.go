package main

import (
	"net/http"
)

/*
	El path que se manda es relativo a la carpeta en la que se encuatran
*/
func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "index.html")
	// })
	/*Permite el acceso a todos los acrivo desde la carpta en la cual
	esta el script*/
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println(r.URL.Path[1:])
	// 	http.ServeFile(w, r, r.URL.Path[1:])
	// })
	/*Fija un prefijo para el acceso a los directorios publicos*/
	fileServer := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fileServer))
	http.ListenAndServe(":8000", nil)
}
