package main

import (
	"fmt"
)

type Human struct {
	name string
}

func (this Human) hablar() string {
	return "bla bla bla"
}

type Dummy struct {
	name string
}

type Tutor struct {
	Human
	Dummy
}

func (this Tutor) hablar() string {
	return "Toma datos del tutor y no del humano"
}

func main() {
	tutor := Tutor{Human{name: "Lenin"}, Dummy{name: "Hi"}}
	fmt.Println(tutor.Dummy.name)
	fmt.Println(tutor.Human.name)
	fmt.Println(tutor.Human.hablar())
	fmt.Println(tutor.hablar())
}
