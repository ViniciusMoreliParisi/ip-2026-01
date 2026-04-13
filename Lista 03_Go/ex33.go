package main

import (
	"fmt"
)

func main() {
	fmt.Println("Vamos calcular quocientes e restos de divisões de modo arcaico?")
	fmt.Println("Lembre-se de usar um numerador maior que o denominador :) ")
	fmt.Print("Digite um numerador da divisão: ")
	var num float64
	fmt.Scan(&num)
	fmt.Print("Digite um denominador da divisão: ")
	var den float64
	fmt.Scan(&den)

	var quociente int
	var resto float64 = num
	var aux float64 = resto

	for quociente = 0; den <= resto; quociente++ {
		aux -= den
		resto = aux
	}
	fmt.Printf("O quociente da divisão de %.2f por %.2f é = %d\n", num, den, quociente)
	fmt.Printf("Quanto ao resto da divisão, está é = %.2f\n", resto)

}
