package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Digite quantos números quiser, até quando puder!")
	fmt.Println("Quando quiser parar, digite o número 30.000 (trinta mil)")

	var soma float64
	var cont float64
	var maior float64 = 1.0
	var menor float64 = 30000.0
	var par float64
	var contpar float64 = 0.0
	var n float64

	for n != 30000.0 {

		fmt.Print("Digite um número: ")
		fmt.Scan(&n)
		if n == 30000.0 {
			break
		}

		soma += n
		cont++

		if n > maior {
			maior = n
		}

		if n < menor {
			menor = n
		}

		if math.Mod(n, 2) == 0 {
			par += n
			contpar++
		}

	}

	contimpar := cont - contpar

	fmt.Printf("A soma dos números digitados é: %.2f\n", soma)
	fmt.Printf("A media dos números digitados é: %.2f\n", soma/cont)
	fmt.Printf("O maior número digitado é: %.2f\n", maior)
	fmt.Printf("O menor número digitado é: %.2f\n", menor)
	fmt.Printf("A média dos números pares é: %.2f\n", par/contpar)
	fmt.Printf("A porcentagem dos números ímpares é: %.2f%%\n", contimpar/cont * 100)

}
