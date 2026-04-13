package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Vamos calcular o seno de um ângulo A?")
	fmt.Print("Digite o ângulo A em πrad (apenas um número após o ponto decimal): ")
	var A float64
	fmt.Scan(&A)

	if A >= 0 {
		fmt.Println(A, "π rad =", 180*A, "°\n")
		if A > 2 && A <= 4 {
			A -= 2
		} else if A > 4 && A <= 6 {
			A -= 4
		} else if A > 6 && A <= 6.3 {
			A -= 6
		}

		

		A *= math.Pi

		var seno = A - (A*A*A)/6 + (A*A*A*A*A)/120 - (A*A*A*A*A*A*A)/5040

		fmt.Printf("SenA = %.4f\n ", seno)
	} else {
		fmt.Println("Ângulo fora do intervalo permitido.")
	}
}
