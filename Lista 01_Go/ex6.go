package main

import "fmt"

func main() {

	fmt.Print("Digite quantas temperaturas deseja converter de F° para C°: ")
	var n int
	fmt.Scan(&n)

	var cont int = 0

	for cont < n {

		fmt.Print("Digite a temperatura em fahrenheit, que deseja converter: ")
		var fahrenheit float64
		fmt.Scan(&fahrenheit)
		celsius := (fahrenheit - 32) * 5 / 9
		fmt.Printf("%.2f Fahrenheit equivale a %.2f celsius\n", fahrenheit, celsius)
		cont++
	}
}
