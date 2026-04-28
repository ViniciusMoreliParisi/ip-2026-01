package main

import (
	"fmt"
)

func soma(numeros []int, n int) int {

	if n == 0 {
		return numeros[n]
	} else {
		return numeros[n] + soma(numeros, n-1)
	}

}

func main() {

	array := make([]int, 5)

	fmt.Print("Vamos determinar a soma de 5 números:\n")
	fmt.Print("Disponhas os cinco números, separados por espaço: ")
	fmt.Scan(&array[0])
	fmt.Scan(&array[1])
	fmt.Scan(&array[2])
	fmt.Scan(&array[3])
	fmt.Scan(&array[4])

	fmt.Print("A soma dos termos digitados é: ", soma(array, len(array)-1))

}
