package main

import "fmt"

func main() {
	fmt.Println("Vamos calcular a média e a soma de dois números (entre 50 e 70)")
	fmt.Print("Digite o primeiro número: ")
	var n1 float64
	fmt.Scan(&n1)
	fmt.Print("Digite o segundo número: ")
	var n2 float64
	fmt.Scan(&n2)
	if n1 < 50 || n1 > 70 || n2 < 50 || n2 > 70 {
		fmt.Println("Os números devem estar entre 50 e 70. Por favor, tente novamente.")
		return
	} else {

		var media float64 = (n1 + n2) / 2
		var soma float64 = n1 + n2
		fmt.Printf("A soma dos números %.2f e  %.2f é: %.2f\n", n1, n2, soma)
		fmt.Printf("A media dos números %.2f e  %.2f é: %.2f\n", n1, n2, media)
	}
}
