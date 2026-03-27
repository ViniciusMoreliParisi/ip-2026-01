package main

import "fmt"

func main() {

	fmt.Print("Digite o valor de compra, e te darei o valor de venda: ")
	var n float64
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("O valor de compra é inválido!")
	}

	if n < 10 {
		fmt.Println("O valor de venda é ", n*1.7)
	}

	if n >= 10 && n < 30 {
		fmt.Println("O valor de venda é ", n*1.5)
	}

	if n >= 30 && n < 50 {
		fmt.Println("O valor de venda é ", n*1.4)
	}

	if n >= 50 {
		fmt.Println("O valor de venda é ", n*1.3)

	}
}
