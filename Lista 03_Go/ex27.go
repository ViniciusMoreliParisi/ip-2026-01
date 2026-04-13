package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Vamos calcular o cosseno de X?")
	fmt.Print("Digite o ângulo X em πrad: ")
	var x float64
	fmt.Scan(&x)
	var cos float64 = 1

	if x > 0 {
		fmt.Print("\n", x, " πrad = ", 180*x, "°  |  ")
		if x > 2 {
			x = math.Mod(x, 2)
		}
		x *= math.Pi

		var frac float64
		var num float64 = 1
		var den float64 = 2
		for i := 1; i <= 20; i++ {
			for j := den; j > 0; j-- {
				num *= x
			}

			frac = num / fatorial(den)
			if i%2 == 0 {
				cos += frac
			} else {
				cos -= frac
			}
			den += 2
			num = 1

		}

		fmt.Printf("CosX = %.4f\n", cos)
	} else {
		fmt.Println("Ângulo fora do intervalo permitido.")
	}

	fmt.Println("\nVamos ver a diferença entre o valor do cosseno calculado \ne o valor do cosseno usando a função math.Cos()?")
	fmt.Printf("\nValor do cosseno usando math.Cos(): %.4f", math.Cos(x))
	fmt.Printf("\nDiferença entre os valores: %.4f", cos-math.Cos(x))

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
