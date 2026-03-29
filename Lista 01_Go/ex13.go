package main

import "fmt"

func main() {

	fmt.Print("Digite sua nota: ")
	var nota float64
	fmt.Scan(&nota)

	var conceito string

	if nota >= 9 && nota <= 10 {
		conceito = "A"
	} else if nota >= 7.5 && nota < 9 {
		conceito = "B"
	} else if nota >= 6 && nota < 7.5 {
		conceito = "C"
	} else if nota >= 0 && nota < 6 {
		conceito = "D"
	}

	fmt.Println("Para nota = ", nota, " Conceito = ", conceito)

}
