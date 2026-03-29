package main

import "fmt"

func main() {
	fmt.Println("A fim de determinar o determinante de uma matriz, ")
	fmt.Print("Digite os valores a, b, c e d (separados por espaço): ")
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)

	det := (a * d) - (b * c)
	fmt.Printf("O valor do determinante é: %.2f", det)

}
