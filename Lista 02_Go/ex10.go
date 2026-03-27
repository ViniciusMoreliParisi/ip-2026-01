package main

import "fmt"

func main() {
	fmt.Println("Segundo a tabela, informe, primeiro, seu número de destino.")
	var n int
	fmt.Scan(&n)

	fmt.Println("Agora, informe se sua passagem é de \"ida\" OU \"ida e volta\" (TUDO MINÚSCULO)")
	var tipo string
	fmt.Scan(&tipo)

	if n == 1 && tipo == "ida" {
		fmt.Println("O preço da passagem é R$ 500,00")

		if n == 1 && tipo == "ida e volta" {
			fmt.Println("O preço da passagem é R$ 900,00")

		}
	}
}
