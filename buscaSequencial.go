package main

import (
	"fmt"
)

func main() {

	fmt.Print("Digite um número, e te direi se ele está presente no meu array: ")
	var n int
	fmt.Scan(&n)

	fmt.Println("Posição do número no array: ", buscaSequencial(n))
	fmt.Print("Caso -1, o número não está presente no array.")

}

func buscaSequencial(n int) int {
	var meuArray [186]int
	for i := 0; i <= 185; i++ {
		meuArray[i] = i

	}
	for i := 0; i <= 185; i++ {
		if meuArray[i] == n {
			return i
			break
		}
	}
	return -1
}
