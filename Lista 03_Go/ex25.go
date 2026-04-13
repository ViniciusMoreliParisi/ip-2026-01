package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Vamos calcular o valor do somatório S?")
	fmt.Print("Digite sim, ou não: ")
	var sim string
	fmt.Scan(&sim)

	sim = strings.ToUpper(sim)

	if sim != "SIM" {
		fmt.Println("Ok, até a próxima!")
	} else {
		fmt.Print("S = ")
		var num float64 = 1
		var sqrtden int
		var soma float64

		for sqrtden = 15; sqrtden >= 1; sqrtden-- {
			if sqrtden%2 != 0 {
				soma += num / float64(sqrtden * sqrtden)
				fmt.Print(" + ", num, "/", sqrtden*sqrtden)
			} else if sqrtden%2 == 0 {
				soma -= num / float64(sqrtden * sqrtden)
				fmt.Print(" - ", num, "/", sqrtden*sqrtden)
			}
			num *= 2
		}
		fmt.Printf("\nS = %.4f", soma)
	}
}
