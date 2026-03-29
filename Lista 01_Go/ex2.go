package main

import "fmt"

func main() {

	fmt.Println("Quantos jogos deseja fazer? ")
	var jogos int
	fmt.Scan(&jogos)

	var cont int = 1

	for cont <= jogos {

		fmt.Print("Qual foi o público total? ")
		var publico float64
		fmt.Scan(&publico)

		fmt.Println("Qual foi a porcentagem da categoria Popular? ")
		var porcentP float64
		fmt.Scan(&porcentP)

		fmt.Println("Qual foi a porcentagem da categoria Geral? ")
		var porcentG float64
		fmt.Scan(&porcentG)

		fmt.Println("Qual foi a porcentagem da categoria Arquibancada? ")
		var porcentA float64
		fmt.Scan(&porcentA)

		fmt.Println("Qual foi a porcentagem da categoria Cadeiras? ")
		var porcentC float64
		fmt.Scan(&porcentC)

		var valorP float64 = 1
		var valorG float64 = 5
		var valorA float64 = 10
		var valorC float64 = 20
		var total float64 = (publico * porcentP * valorP) + (publico * porcentG * valorG) + (publico * porcentA * valorA) + (publico * porcentC * valorC)

		fmt.Printf("A RENDA DO JOGO N.%d É %2.f", cont, total)

		cont++
	}
}
