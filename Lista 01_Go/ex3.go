package main

import "fmt"

func main() {

	fmt.Print("Digite três algarismos, separados por espaço. Estes serão, respectivamente, a centena, dezena e unidade: ")
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a > 9 || b > 9 || c > 9 {
		fmt.Print("Algum algarismo é inválido \nmaior que 9*")
	} else {
		a = 100 * a
		b = 10 * b
		c = c
		var numero int = a + b + c
		var numero2 int = numero * numero
		fmt.Print(numero, ", ", numero2)
	}

}
