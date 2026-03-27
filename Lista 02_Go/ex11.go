package main

import "fmt"

func main() {

	fmt.Print("Digite o valor de x, e calcularei f(x) = 8 / (2-x): ")
	var x float64
	fmt.Scan(&x)

	var y float64 = 8 / (2 - x)

	fmt.Printf("O valor de f(x) é = %.2f ", y)

}
