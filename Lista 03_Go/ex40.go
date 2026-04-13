package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Vamos calcular as vendas e o lucro esperado para os preços dos ingressos?")
	fmt.Print("Digite sim, ou não: ")
	var sim string
	fmt.Scan(&sim)

	sim = strings.ToUpper(sim)

	if sim != "SIM" {
		fmt.Println("Ok, até a próxima!")
	} else {

		var vendas float64 = 130.00
		var despesa float64 = 300.00
		var barateamento float64 = 1.00

		for preçoIngresso := 6.00; preçoIngresso >= 1; preçoIngresso -= 0.6 {

			var receita float64 = preçoIngresso * vendas
			var lucro float64 = receita - despesa

			fmt.Printf("\n\n%d) PREÇO DO INGRESSO = R$%.2f   | INGRESSOS VENDIDOS = %.2f", int(barateamento), preçoIngresso, vendas)
			fmt.Printf("\n   LUCRO MÁXIMO      = R$%.2f", lucro)

			vendas += 30 * barateamento
			receita = preçoIngresso * vendas
			lucro = receita - despesa
			barateamento += 1.00
		}

	}
}
