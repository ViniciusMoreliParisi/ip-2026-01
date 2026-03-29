package main

import "fmt"

func main() {

	fmt.Println("Digite dois números (separados por espaço)\nO primeiro deve ser par (é a partir de onde o contador começará)\nO seguno é a quantidade de termos da sequência de números pares :) ")
	var a, b int
	fmt.Scan(&a, &b)

	if a%2 != 0 {
		fmt.Print("Amigo, o primeiro número deve ser parrrr ._.")
	} else {
		for b > 0 {
			fmt.Print(a, " ")
			a = a + 2
			b--

		}
	}
}
