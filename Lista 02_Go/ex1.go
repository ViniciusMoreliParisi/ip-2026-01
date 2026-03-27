package main

import "fmt"

func main() {

	fmt.Println("Digite um número, e te direi se ele é par ou ímpar.")
	var n int
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println("Esse número é par.")
	} else {
		fmt.Println("Esse número é ímpar.")
	}
}
