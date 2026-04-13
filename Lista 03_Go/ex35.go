package main

import (
	"fmt"
)

func main() {
	fmt.Println("Vamos transformar um número base 10 em um número binário!")
	fmt.Print("Digite o nosso número que será formatado: ")
	var n int
	fmt.Scan(&n)

	fmt.Printf("O número %.d em binário é: %b ", n, n)

}
