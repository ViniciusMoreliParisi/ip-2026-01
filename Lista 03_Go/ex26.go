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

		var num float64 = 100
		var den float64 = 0
		var soma float64

		for i := 1; i <= 20; i++ {

			soma += num / fatorial(den)

			num--
			den++
		}
		fmt.Printf("O valor do somatório S é: %.2f", soma)
	}
}

func fatorial(n float64) float64 {

	if n <= 1 {
		return 1
	} else {
		for i := n - 1; i >= 1; i-- {
			n = n * i
		}
		return n
	}
}
