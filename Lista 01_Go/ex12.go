package main

import "fmt"

func main() {

	fmt.Print("Digite as horas de uso da charrete: ")
	var n int
	fmt.Scan(&n)

	gasto := (n-n%3)/3*10 + (n%3)*5

	fmt.Print("O valor a pagar é: ", gasto)

}
