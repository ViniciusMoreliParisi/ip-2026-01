package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Vamos calcular o volume de uma pirâmide de base hexagonal? :)")
	fmt.Print("Digite, primeiro, o valor da sua altura: ")
	var h float64
	fmt.Scan(&h)

	fmt.Print("Digite, agora, o valor de uma das arestas do hexágono regular: ")
	var a float64
	fmt.Scan(&a)

	var Ab float64
	Ab = 3 * a * a * math.Sqrt(3) / 2

	var volume float64
	volume = 3 * Ab * h / 3

	fmt.Printf("O volume da pirâmede é %.2f metros cúbicos! ._.", volume)

}
