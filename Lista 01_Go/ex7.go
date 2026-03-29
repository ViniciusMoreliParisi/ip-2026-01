package main

import "fmt"

func main() {

	fmt.Print("Digite a temperatura em fahrenheit, que deseja converter: ")
	var fahrenheit float64
	fmt.Scan(&fahrenheit)
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f Fahrenheit equivale a %.2f celsius\n", fahrenheit, celsius)

	fmt.Print("Determine a quantidade de polegadas que deseja converter em mm: ")
	var polegadas float64
	fmt.Scan(&polegadas)
	mm := (polegadas * 25.4)
	fmt.Printf("A quantidade de chuva em mm é: %.2f", mm)
}
