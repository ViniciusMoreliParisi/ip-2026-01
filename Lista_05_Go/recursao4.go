package main

import (
	"fmt"
)

func main() {
	fmt.Println("Vamos fazer uma conversão decimal -> binário")
	fmt.Print("Digite um número (1 a 15) a ser convertido: ")
	var n int
	fmt.Scan(&n)

	array := make([]int, 5)

	fmt.Print("Seu número (", n, "), em binário, é: ", binar(n, 4, array))

}

func binar(n, index int, binario []int) []int {
	if n > 0 {
		binario[index] = n % 2
		n = (n - n%2) / 2
		return binar(n, index-1, binario)

	}
	return binario
}
