package main

import "fmt"

func main() {

	fmt.Print("Digite um número, e te direi se ele é divisível por 5, ou não: ")
	var n int
	fmt.Scan(&n)
	if n%5 == 0 {
		fmt.Println("O número é divisível por 5.")
	} else {
		fmt.Println("O número não é divisível por 5.")
	}
}
