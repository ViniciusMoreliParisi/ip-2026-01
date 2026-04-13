package main

import (
	"fmt"
)

func main() {
	fmt.Println("Vamos transformar um número base 10 em um número base 16!")
	fmt.Print("Digite o nosso número que será formatado: ")
	var n int
	fmt.Scan(&n)

	fmt.Printf("O número %.d em hexadecimal é: %x ", n, n)

}
