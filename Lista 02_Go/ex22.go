package main

import "fmt"

func main() {
	fmt.Print("Olá, colaborador! Informe sua matrícula -> ")
	var matricula int
	fmt.Scan(&matricula)

	fmt.Print("Nos informe a quantidade de horas-extras: ")
	var horasExtras float64
	fmt.Scan(&horasExtras)

	salarioMinimo := 788.00
	valorhoraextra := 10.00

	salarioBruto := (salarioMinimo * 3) + (horasExtras * valorhoraextra)

	var salarioLiquido float64

	if salarioBruto > 1500 && salarioBruto <= 2000 {
		salarioLiquido = salarioBruto * 0.88
	} else {
		salarioLiquido = salarioBruto * 0.68
	}

	fmt.Println("\nColaborador, de matrícula ", matricula)
	fmt.Printf("Seu salário, desse mês, é: R$%2.f", salarioLiquido)
}
