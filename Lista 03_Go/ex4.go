package main

import "fmt"

func main() {
	fmt.Println("Digite quadrados perfeitos, até quando puder!")

	var n int
	var prova int
	var i int = 1
	var ehquadrado bool = true

	for ehquadrado == true {
		fmt.Print("Digite o número: ")
		fmt.Scan(&n)
		if n < 0 || n == 0 {
			fmt.Print("Digite um número positivo e diferente de zero, por favor.\n")
		} else {
			for prova < n {
				prova = i * i
				i++
			}
		}

		if prova != n {
			ehquadrado = false
			fmt.Println("O número", n, "não é um quadrado perfeito.")
		}

	}

}
