package main

import "fmt"

func main() {
	fmt.Print("Para calcular o valor de delta, digite os coeficientes a, b e c, separados por espaços: ")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	delta := (b * b) - (4 * a * c)
	fmt.Printf("O valor de delta é: %.2f\n", delta)

}
