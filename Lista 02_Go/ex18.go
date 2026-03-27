package main

import "fmt"

func main() {

	fmt.Println("DIAS DA SEMANA: \n- Domingo (1) \n- Segunda (2) \n- Terça (3) \n- Quarta (4) \n- Quinta (5) \n- Sexta (6) \n- Sábado (7)")
	fmt.Print("Digite o número correpsondente ao dia de hoje: ")
	var dia int
	fmt.Scan(&dia)

	fmt.Print("Digite o preço do filme: ")
	var preçoInicial float64
	fmt.Scan(&preçoInicial)

	var preçoFinal float64

	switch dia {
	case 2, 3, 5:
		preçoFinal = preçoInicial * 0.6
		fmt.Println("Desconto de 40% aplicado!")

	case 4, 6, 7, 1:
		preçoFinal = preçoInicial

		break
	default:
		fmt.Println("Número do dia inválido! Por favor, escolha um número entre 1 e 7.")
		return
	}

	fmt.Println("O filme está na sessão \"Lançamentos\"? (S/N) ")
	var lançamento string
	fmt.Scan(&lançamento)

	switch lançamento {
	case "S", "s":
		preçoFinal = preçoFinal * 1.15
		fmt.Println("Acréscimo de 15% aplicado por ser um lançamento :)")
	case "N", "n":
		preçoFinal = preçoFinal
		break
	default:
		fmt.Println("Opção inválida!")
		return
	}

	fmt.Printf("O preço final do filme alugado é: R$%.2f", preçoFinal)
}
