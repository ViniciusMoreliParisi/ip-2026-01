package main

import "fmt"

func main() {

	var soma int

	for i := 1; i <= 20; i++ {
		fmt.Print(i, " | ")
		soma += i
	}

	fmt.Print("\n\n A soma dos números de 1 a 20 é: ", soma)

}
