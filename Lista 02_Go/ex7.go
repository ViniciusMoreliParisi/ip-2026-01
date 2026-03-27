package main

import "fmt"

func main() {

	fmt.Print("Digite três números, e eu os colocarei em ordem crescente: ")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	var menor float64
	var meio float64
	var maior float64

	if a < b && a < c {
		menor = a
	}
	if b < a && b < c {
		menor = b
	}
	if c < a && c < b {
		menor = c
	}

	if a > b && a > c {
		maior = a
	}
	if b > a && b > c {
		maior = b
	}
	if c > a && c > b {
		maior = c
	}

	if a != menor && a != maior {
		meio = a
	}

	if b != menor && b != maior {
		meio = b
	}

	if c != menor && c != maior {
		meio = c
	}

	fmt.Println("A ordem crescente dos números é: ", menor, meio, maior)

}
