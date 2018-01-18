package main

import "fmt"

/*
	== igual a
	!= diferente de
	< menor que
	> mayor que
	>= mayor igual que
	<= menor igual que
	&& and
	|| or
*/
func main() {
	x := 10
	y := 12
	if x >= y {
		fmt.Printf("%d es mayor que %d\n", x, y)
	} else if x < y {
		fmt.Println("Entre al else if")
	} else {
		fmt.Println("Entre else")
	}
}
