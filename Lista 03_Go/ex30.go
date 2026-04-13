package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Bora calcular o volume de uma esfera bem maneira?")
	fmt.Print("Digite o raio da esfera (múltiplos de 0,5): ")
	var r float64
	fmt.Scan(&r)

	var volesfera float64

	if r < 0 || r > 20 {
		fmt.Print("O raio da esfera é INVÁLIDO! (< 0, ou > 20)")
	} else {
		volesfera = 4.0 / 3.0 * math.Pi * math.Pow(r, 3)
		fmt.Printf("\nO volume da esfera é: %.2f cm³", volesfera)
	}
}
