package main

import "fmt"

func main() {
	fmt.Println("Vamos criar uma série de Fetuccine?")
	fmt.Print("Digite o primeiro termo: ")
	var n1 float64
	fmt.Scan(&n1)
	fmt.Print("Digite o segundo termo: ")
	var n2 float64
	fmt.Scan(&n2)

	fmt.Print("Digite o número de termos > 3: ")
	var termos int
	fmt.Scan(&termos)

	var proximo float64

	fmt.Print("Série de Fetuccine: ", n1, ", ", n2, ", ")

	for i := 1; i <= termos-2; i++ {

		if i%2 != 0 {
			proximo = n2 + n1
			fmt.Print(proximo)
		}

		if i%2 == 0 {
			proximo = n2 - n1
			fmt.Print(proximo)
		}

		if i < termos - 2 {
			fmt.Print(", ")
		}

		n1 = n2
		n2 = proximo

	}
}
