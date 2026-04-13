package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Determine a idade, altura e o peso de quantas pessoas quiser, até quando puder!")

	cont := 1
	contmais50 := 0.0
	altura10e20 := 0.0
	cont10e20 := 0.0
	contmenos40kg := 0.0

	for fim := "continue"; fim != "FIM"; {

		fmt.Printf("Digite a idade da %d° pessoa : ", cont)
		var idade int
		fmt.Scan(&idade)
		fmt.Printf("Digite a altura da %d° pessoa : ", cont)
		var altura float64
		fmt.Scan(&altura)
		fmt.Printf("Digite o peso da %d° pessoa : ", cont)
		var peso float64
		fmt.Scan(&peso)

		cont++

		if idade > 50 {
			contmais50++
		} else if idade > 10 && idade < 20 {
			altura10e20 += altura
			cont10e20++
		}

		if peso < 40 {
			contmenos40kg++
		}

		fmt.Print("Digite fim, se terminou de computar os dados: ")
		fmt.Scan(&fim)
		fim = strings.ToUpper(fim)

	}

	fmt.Println("A quantidade de pessoas com idade superior a 50 anos é ", contmais50)
	if cont10e20 > 0 {
		fmt.Println("Média das alturas das pessoas com idade entre 10 e 20 anos: ", altura10e20/float64(cont10e20))
	} else {
		fmt.Println("Não houve nenhuma pessoa com idade entre 10 e 20 anos.")
	}
	fmt.Println("Porcentagem de pessoas com peso inferior a 40 quilos: ", float64(contmenos40kg)/float64(cont)*100, "%")

}
