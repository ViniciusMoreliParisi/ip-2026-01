package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("\nFaça um programa que calcule e escreva o número de grãos de milho que se pode colocar em um tabuleiro de")
	fmt.Println("xadrez, colocando 1 no primeiro quadro e nos quadros seguintes o dobro do quadro anterior. São 64 quadros no")
	fmt.Println("total.")
	fmt.Println("\n--------------------------------------------------------------------------------------------------------------")

	var graos_por_casa float64 = 1
	var total_de_graos float64

	fmt.Println("\nBora calcular esses graos de milho?")
	fmt.Print("Digite sim, ou não: ")
	var sim string
	fmt.Scan(&sim)

	sim = strings.ToUpper(sim)

	if sim != "SIM" {
		fmt.Println("Ok, até a próxima!")
	} else {

		for i := 1; i <= 64; i++ {
			total_de_graos += graos_por_casa
			graos_por_casa *= 2
		}
	}
	fmt.Printf("\nPode-se colocar, em um tabuleiro de xadrez, %0.f grãos de milho.", total_de_graos)
}
