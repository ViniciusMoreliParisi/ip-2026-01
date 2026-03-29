package main

import "fmt"

func main() {

	fmt.Print("Insira o seu salário, para saber o seu valor reajustado: ")
	var salario float64
	fmt.Scan(&salario)

	switch {
	case salario <= 300:
		fmt.Printf("Seu salário reajustado é: R$%.2f", salario*1.5)

	case salario > 300:
		fmt.Printf("Seu salário reajustado é: R$%.2f", salario*1.3)

	}
}
