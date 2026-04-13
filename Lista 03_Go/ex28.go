package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("Bora calcular o valor de ∏ com 51 termos?")
	fmt.Print("Digite sim, ou não: ")
	var sim string
	fmt.Scan(&sim)

	sim = strings.ToUpper(sim)

	var num float64 = 1
	var den float64 = 1
	var frac float64
	var S float64
	var incógnita float64

	if sim != "SIM" {
		fmt.Println("Ok, até a próxima!")
	} else {

		for i := 1; i <= 51; i++ {
			frac = num / math.Pow(den, 3)
			den += 2
			if i%2 == 0 {
				S -= frac
			} else {
				S += frac
			}
		}

		incógnita = math.Cbrt(32 * S)
		fmt.Printf("O valor de ∏ é, aproximadamente: %.6f\n", incógnita)
	}
}
