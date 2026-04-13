package main

import "fmt"

func main() {
	fmt.Println("Vamos montar a sequência de Fibonacci?")
	fmt.Print("Digite o número de termos da sequência: ")
	var n float64
	fmt.Scan(&n)

	if n < 3 {
		fmt.Println("A sequência de Fibonacci deve ter pelo menos 3 termos.")
		return
	} else {

		var a int = 0
		var b int = 1

		for i := n; i > 0; i-- {

			fmt.Print(a)

			if i > 1 {
				fmt.Print(" - ")

				aux := a
				a = b
				b = aux + b

			}

		}
	}
}
