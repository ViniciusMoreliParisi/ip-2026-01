package main

import "fmt"

func main() {
	fmt.Println("A fim de calcularmos seu gasto com energia,")
	fmt.Print("Digite o valor do seu sálario: ")
	var salario float64
	fmt.Scan(&salario)

	fmt.Print("Digite a quantidade de kW gastos: ")
	var kW float64
	fmt.Scan(&kW)

	var valorkW float64 = 0.007 * salario

	fmt.Printf("O custo do kW é: R$%.2f\n", valorkW)
	fmt.Printf("O custo do consumo é: R$%.2f\n", (valorkW * kW))
	fmt.Printf("O custo com desconto é: R$%.2f\n", (valorkW * kW * 0.9))
}
