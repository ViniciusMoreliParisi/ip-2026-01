package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Digite um número: ")
	var n float64
	fmt.Scan(&n)

	var a float64
	var b float64

	a = math.Sqrt(n)
	b = n * n
	if n >= 0 {
		fmt.Println("A raiz qudrada desse número é: ", a)
	} else {
		fmt.Println("O quadrado desse número é: ", b)

	}
}
