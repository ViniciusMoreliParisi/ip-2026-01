package main

import (
	"fmt"
)

func main() {
	fmt.Println("Vamos inverter a ordem da sua lista")
	fmt.Print("Digite a quantidade de termos da sua lista: ")
	var n int
	fmt.Scan(&n)

	array := make([]int, n)

	fmt.Print("Declare sua lista: ")

	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	fmt.Println("\nA lista original é: ", array)

	fmt.Print("Sua lista invertida é: ", invert(array, 0, n-1))

}

func invert(lista1 []int, esq, dir int) []int {
	if esq > dir {
		return lista1
	}

	lista1[esq], lista1[dir] = lista1[dir], lista1[esq]

	return invert(lista1, esq+1, dir-1)

}
