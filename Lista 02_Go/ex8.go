package main

import "fmt"

func main() {

	fmt.Print("Digite um número, e te direi se ele está entre 20 e 90: ")
	var n float64
	fmt.Scan(&n)

	if 90 > n && n > 20 {
		fmt.Println("O número está entre 20 e 90.")
	} else {
		fmt.Println("O número não está entre 20 e 90.")
	}

}
