package main

import "fmt"

func main() {
	fmt.Print("Para começar, digite o valor do produto: ")
	var produto float64
	fmt.Scan(&produto)

	fmt.Println("Agora, determine a forma de pagamento: \nÀ vista, dinheiro ou cheque (1) \nCartão de crédito (2): ")
	var pagamento float64
	fmt.Scan(&pagamento)

	var valorFinal float64

	switch pagamento {
	case 1:
		valorFinal = produto * 0.9

	case 2:
		fmt.Print("Você escolheu a opção de crédito \nDigite o número de parcelas (1, 2 ou 3): ")
		var parcelas float64
		fmt.Scan(&parcelas)
		switch parcelas {
		case 1:
			valorFinal = produto * 0.95
		case 2:
			valorFinal = produto
		case 3:
			valorFinal = produto * 1.1

			break
		default:
			fmt.Println("Número de parcelas inválido! Por favor, escolha 1, 2 ou 3.")
			return

		}
	}
	fmt.Println("O seu preço final é R$:", valorFinal)
}
