package main

import (
	"fmt"
)

func main() {

	var cpf [11]int

	fmt.Println("Vamos validar o seu CPF!")
	fmt.Print("Digite o seu CPF (separados por espaços): ")
	for i := 0; i < 11; i++ {
		fmt.Scan(&cpf[i])
	}

	var soma int
	var j int = 8

	for i := 2; i <= 10; i++ {

		soma += i * cpf[j]
		j--
	}

	var dig1 int

	if soma%11 < 2 {
		dig1 = 0
	} else if soma%11 >= 2 {
		dig1 = 11 - (soma % 11)
	}

	if cpf[9] != dig1 {
		fmt.Println("O primeiro dígito verificador está incorreto.")
		fmt.Println("Portanto, seu CPF é inválido.")

	} else {
		fmt.Println("O primeiro dígito verificador está correto.")

		soma = 0
		j = 9

		for i := 2; i <= 11; i++ {
			soma += i * cpf[j]
			j--
		}

		var dig2 int

		if soma%11 < 2 {
			dig2 = 0
		} else if soma%11 >= 2 {
			dig2 = 11 - (soma % 11)
		}
        
		
		if cpf[10] != dig2{
			fmt.Println("O segundo dígito verificador está incorreto.")
			fmt.Println("Portanto, seu CPF é inválido.")
		} else {
			fmt.Println("O segundo dígito verificador está correto.")
			fmt.Println("Parabéns! Seu CPF é válido.")
		}

		for i := 0; i < 11; i++ {

			if i == 0 {
				fmt.Print("CPF válido: ")
			}

			if i == 2 || i == 5 {
				fmt.Print(cpf[i], ".")
			} else if i == 8 {
				fmt.Print(cpf[i], "-")
			} else {
				fmt.Print(cpf[i])
			}

		}
	}
}
