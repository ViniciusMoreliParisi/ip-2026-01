package main

import (
	"fmt"
)

func main() {
	fmt.Println("Vamos descobrir o mínimo múltiplo comum de dois números!")
	fmt.Print("Digite o primeiro número: ")
	var n1 int
	fmt.Scan(&n1)
	fmt.Print("Digite o segundo número: ")
	var n2 int
	fmt.Scan(&n2)

	var mmc1 bool = false
	var mmc2 bool = false

	var i int = 2

	for mmc1 == false || mmc2 == false {

		mmc1 = false
		mmc2 = false

		if n1%i == 0 {
			mmc1 = true
		}
		if n2%i == 0 {
			mmc2 = true
		}

		i++

		if mmc1 == true && mmc2 == true {
			fmt.Printf("O mínimo múltiplo comum entre %d e %d é: %d", n1, n2, i-1)
			break
		} else if i > n1 || i > n2 {
			fmt.Println("Não foi possível encontrar um  mínimo múltiplo comum para os números digitados.")
			break
		}
	}
}
