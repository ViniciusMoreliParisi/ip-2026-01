package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Vamos efetuar uma soma de P.A?")
	fmt.Print("Digite sim, ou não: ")
	var sim string
	fmt.Scan(&sim)

	sim = strings.ToUpper(sim)

	if sim != "SIM" {
		fmt.Println("Ok, até a próxima!")
	} else {

		var soma float64
		var numerador float64 = 38.00
		var denominador float64 = 1.00

		for i := numerador; i >= 2; i-- {
			soma += float64(numerador*(numerador-1.00)) / float64(denominador)

			fmt.Print(numerador, "*", (numerador - 1.00), " /", denominador, "\n")

			numerador--
			denominador++
		}
		fmt.Printf("A soma é: %.2f\n", soma)
	}
}
