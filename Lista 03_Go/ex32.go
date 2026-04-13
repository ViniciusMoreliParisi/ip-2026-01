package main

import (
	"fmt"
)

func main() {
	fmt.Println("Vamos multiplicar dois números à moda antiga!")
	fmt.Print("Digite um primeiro número: ")
	var n1 int
	fmt.Scan(&n1)
	fmt.Print("Digite um segundo número: ")
	var n2 int
	fmt.Scan(&n2)

	var soma int = n1

	for i := 1; i < n2; i++ {
		fmt.Printf("%d° Soma: %d + %d = %d\n", i, soma, n1, soma+n1)
		soma += n1

	}
	fmt.Printf("O resultado da multiplicação é: %d", soma)
}
