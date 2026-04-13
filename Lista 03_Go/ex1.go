package main

import "fmt"

func main() {
	fmt.Println("Vamos calcular uma potência bem maneira?")
	fmt.Print("Digite o número que será a base: ")
	var n float64
	fmt.Scan(&n)
	fmt.Print("Digite, agora, seu expoente: ")
	var e int
	fmt.Scan(&e)

	var resposta float64 = 1

	for cont := e; cont > 0; cont-- {
		resposta *= n
	}

	fmt.Printf("O resultado de %.2f elevado a %d é: %.2f\n", n, e, resposta)

}
