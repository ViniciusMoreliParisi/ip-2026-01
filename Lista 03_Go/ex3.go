package main

import "fmt"

func main() {
	fmt.Print("Carlos, forneça o seu salário: ")
	var salc float64
	fmt.Scan(&salc)

	salj := salc / 3.00

	fmt.Printf("\nJoão, seu salário é: %.2f \n", salj)

	fmt.Println("Em quanto tempo de aplicação seus investimentos vão se equivaler?")

	var investc float64 = salc
	var investj float64 = salj

	var cont int

	for investc > investj {

		investc *= 1.02
		investj *= 1.05

		cont++
	}

	fmt.Printf("Ao final de %d meses, os investimentos de Carlos e João vão se equivaler.\n", cont)
	fmt.Printf("O valor do investimento de Carlos será %.2f\n", investc)
	fmt.Printf("O valor do investimento de João será %.2f\n", investj)
}
