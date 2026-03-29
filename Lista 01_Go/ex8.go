package main

import "fmt"

func main() {

	fmt.Print("Digite o raio da latinha: ")
	var r float64
	fmt.Scan(&r)

	var h float64
	fmt.Print("Digite a altura da latinha: ")
	fmt.Scan(&h)

	area := ((2 * 3.14 * r * h) + (2 * 3.14 * r * r))
	fmt.Printf("Área da latinha: %.2f\n", area)
	fmt.Printf("O valor da latinha é: R$%.2f", area*100)
}
