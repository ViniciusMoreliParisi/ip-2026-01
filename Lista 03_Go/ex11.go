package main

import "fmt"

func main() {
	fmt.Println("Dê-me um número e calcularei o seu fatorial!")
	fmt.Print("Digite o número: ")
	var n int
	fmt.Scan(&n)
    aux := n

	if n <= 1 {
		fmt.Print("O fatorial de", n, " é = 1")
	} else if n > 1 {
		var fatorial int = n
		for i := n; i > 1; i-- {
			fatorial *= (n - 1)
			n--
		}
		fmt.Print("O fatorial de ", aux, " é = ", fatorial)

	}

}
