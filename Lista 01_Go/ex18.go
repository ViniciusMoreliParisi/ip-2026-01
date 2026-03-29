package main

import "fmt"

func main() {

	fmt.Println("Vamos calcular a soma dos termos de uma P.A (Progressão Aritmética)? ")
	fmt.Print("Digite os valor do a1, r e n: ")
	var a1, r, n int
	fmt.Scan(&a1, &r, &n)

	var soma int

	for n > 0 {
		soma += r
		n--
	}

	fmt.Println("A soma dos termos da P.A é ", soma)

}
