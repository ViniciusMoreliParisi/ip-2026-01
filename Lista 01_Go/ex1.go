package main

import "fmt"

func main() {

	fmt.Print("Digite suas três notas, separadas por espaço: ")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	media := (a + b + c) / 3

	var situação string

	if media >= 6 {
		situação = "APROVADO!"
	} else {
		situação = "DESAPROVADO!"
	}

	fmt.Printf("MEDIA = %2.f\n", media)
	fmt.Print(situação, "\n")
}
