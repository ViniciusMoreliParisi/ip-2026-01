package main

import (
	"fmt"
)

func main() {

	var cpf [11]int

	fmt.Println("Vamos calcular os dígitos verificadores do seu CPF!")
	fmt.Print("Digite os primeiros 9 dígitos do seu CPF (separados por espaços): ")
	for i := 0; i < 9; i++ {
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

	cpf[9] = dig1
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

	cpf[10] = dig2

	for i := 0; i < 11; i++ {

		if i == 0 {
			fmt.Print("O CPF completo é: ")
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
