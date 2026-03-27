package main

import "fmt"

func main() {

	fmt.Println("Digite um número, e te direi se ele é positivo, negativo, ou zero.")
	var n int
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("O número é zero.")
	} else if n > 0 {
		fmt.Println("O número é positivo.")
	} else {
		fmt.Println("O número é negativo.")
	}
}
