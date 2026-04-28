package main

import (
	"fmt"
)

func main() {

	fmt.Println("Vamos calcular a expoenenciação de n, na base a?")
	fmt.Print("Digite o valor de a: ")
	var a int
	fmt.Scan(&a)
	fmt.Print("Digite o valor de n: ")
	var n int
	fmt.Scan(&n)

	fmt.Println("O resultado é: ", exponencial(a, n))

}

func exponencial(a, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return a
	}
	return a * exponencial(a, n-1)
}
