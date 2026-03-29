package main

import (
	"fmt"
)

func main() {

	fmt.Print("Digite um número: ")
	var n int
	fmt.Scan(&n)

	var par int = 2
	var par2 int
	cont := n

	if n%2 != 0 {
		cont--
	}

	for cont > 0 {

		par2 = par * par
		fmt.Println(par, "*", par, "=", par2)

		par = par + 2
		cont = cont - 2
	}
}
