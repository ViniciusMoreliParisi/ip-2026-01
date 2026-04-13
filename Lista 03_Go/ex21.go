package main

import "fmt"

func main() {
	fmt.Println("Vamos calcular B elevado a n?")
	fmt.Print("Digite o valor de B: ")
	var B float64
	fmt.Scan(&B)
	fmt.Print("Digite, agora, o valor de n: ")
	var n int
	fmt.Scan(&n)

	var resposta float64 = 1

	if B >= 2 && n > 1 {
		for cont := n; cont > 0; cont-- {
			resposta *= B
		}
	fmt.Printf("O resultado de %.2f elevado a %d é: %.2f\n", B, n, resposta)
    }else{
    fmt.Println("Valor de B deve ser maior ou igual a 2 e n deve ser maior que 1.")
}



	

}
