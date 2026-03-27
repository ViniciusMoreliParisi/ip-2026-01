package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("CALCULADORA BHÁSKARA! ;) \n================================")
	fmt.Print("Lembrando que a equação é do tipo ax² + bx + c = 0, \ndigite os coeficientes a, b e c, separados por espaços: ")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	delta := (b * b) - (4 * a * c)

	if delta == 0 {
		x := -b / (2 * a)
		fmt.Println("A equação tem RAIZ ÚNICA = ", x)
	} else if delta > 0 {
		raizDelta := math.Sqrt(delta)
		x1 := (-b + raizDelta) / (2 * a)
		x2 := (-b - raizDelta) / (2 * a)
		fmt.Println("A equação possúi RAÍZES DISTINTAS =", x1, "e ", x2)
	} else if delta < 0 {
		fmt.Println("A equação tem RAIZ IMAGINÁRIA, pois delta < 0")
	}

}
