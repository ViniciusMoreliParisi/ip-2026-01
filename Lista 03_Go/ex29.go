package main

import "fmt"

func main() {
	fmt.Println("Vamos somar todos os números naturais até um número que você escolher.")
	fmt.Print("Digite um número: ")
	var n int
	fmt.Scan(&n)

	var soma int

	for i := 1; i <= n; i++ {
		soma += i
	}

	fmt.Printf("A soma dos %d termos é = %d", n, soma)

}
