package main

import (
	"encoding/json"
	"net/http"
)

/*
	atributos publicos Primera letra mayuscula
	atributos privados primera letra minuscula
*/
type Curso struct {
	Title        string `json:"titulo"`
	NumeroVideos int    `json:"numero_videos"`
}

type Cursos []Curso

func main() {
	cursos := Cursos{
		Curso{Title: "Curso de go1", NumeroVideos: 30},
		Curso{Title: "Curso de go2", NumeroVideos: 30},
		Curso{Title: "Curso de go3", NumeroVideos: 30},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(cursos)
	})
	http.ListenAndServe(":8000", nil)
}
