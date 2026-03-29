package main

import "fmt"

func main() {

	fmt.Println("CALCULADORA SANEAGO \n ==========================")
	fmt.Println("Digite seu tipo de consumo \n- Residencial (01) \n- Comercial   (02) \n- Industrial  (03) ")

	var conta int
	fmt.Scan(&conta)

	fmt.Print("Agora, digite o valor do consumo em m³: ")
	var consumo float64
	fmt.Scan(&consumo)

	var gasto float64

	switch conta {
	case 1:
		gasto = 5 + consumo*0.05
	case 2:
		if consumo <= 80 {
			gasto = 500
		} else if consumo > 80 {
			gasto = 500 + (consumo-80)*0.25
		}
	case 3:
		if consumo <= 10 {
			gasto = 800
		} else if consumo > 100 {
			gasto = 800 + (consumo-100)*0.04
		}

	default:
		fmt.Println("Tipo de conta inválida. Por favor, escolha entre 01, 02 ou 03.")
		return
	}

	fmt.Printf("Sua conta, do tipo %d, com consumo de %.2f m³, é de R$%.2f", conta, consumo, gasto)
}
