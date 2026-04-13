package main

import "fmt"

func main() {
	fmt.Println("Escreva um número traingular, e irei validá-lo para você!!")
	fmt.Print("Digite o número: ")
	var n int
	fmt.Scan(&n)

	var numtriangular int
	var i int

	for i = 1; numtriangular < n; i++ {

		numtriangular = i * (i + 1) * (i + 2)
	}

	if numtriangular == n {
		fmt.Print("O número ", n, " é um número triangular :)")
	} else {
		fmt.Print("O número ", n, " não é um número triangular :(")
	}
}
