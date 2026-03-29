package main

import "fmt"

func main() {

	fmt.Print("Digite um número inteiro e positivo: ")
	var n float64
	fmt.Scan(&n)

	var soma float64
	cont := n
	n = 1

	for cont > 0 {

		soma += 1.0 / n
		n++
		cont--
	}

	fmt.Printf("%.6f", soma)
}
