package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Vamos calcular áreas e volumes?")
	fmt.Println("Escolha qual sólido geométrico deseja calcular:\n1 - Cone\n2 - Cilindro Reto\n3 - Esfera")
	fmt.Print("Digite o número correspondente ao sólido geométrico: ")
	var solido int
	fmt.Scanln(&solido)

	var r, h float64

	switch solido {
	case 1, 2:
		fmt.Print("Digite o raio da base (inteiro):")
		fmt.Scanln(&r)
		fmt.Print("Digite a altura (inteiro): ")
		fmt.Scanln(&h)
	case 3:
		fmt.Print("Digite o raio da esfera (inteiro):")
		fmt.Scanln(&r)
		break
	default:
		fmt.Println("Opção inválida. Por favor, escolha 1, 2 ou 3.")
		return
	}

	var area float64
	var volume float64

	switch solido {
	case 1:
		volume = 3.14 * r * r * h / 3
		area = 3.14 * r * math.Sqrt(r*r+h*h)
	case 2:
		volume = 3.14 * r * r * h
		area = (2 * 3.14 * r * h) + (2 * 3.14 * r * r)
	case 3:
		volume = 4 / 3 * 3.14 * r * r * r
		area = 4 * 3.14 * r * r
		break
	}

	fmt.Printf("A área do sólido escolhido é = %.2f", area)
	fmt.Printf("O volume do sólido escolhido é = %.2f", volume)

}
