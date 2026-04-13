package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Vamos calcular o resultado da série de N termos?")
	fmt.Print("Digite sim, ou não: ")
	var sim string
	fmt.Scan(&sim)

	sim = strings.ToUpper(sim)

	if sim != "SIM" {
		fmt.Println("Ok, até a próxima!")
	} else {
		fmt.Print("Digite o valor de N: ")
		var n int
		fmt.Scan(&n)

		var num float64 = 1000
		var den float64 = 1
		var soma float64

		for i := 1; i <= n; i++ {

			if i%2 != 0 {
				soma += float64(num) / float64(den)
				fmt.Print("+ ", num, "/", den, "\n")

			} else if i%2 == 0 {
				soma -= float64(num) / float64(den)
				fmt.Print("- ", num, "/", den, "\n")
			}
			num -= 3
			den += 1
		}

		fmt.Printf("\nSOMA = %.2f", soma)
	}

}
