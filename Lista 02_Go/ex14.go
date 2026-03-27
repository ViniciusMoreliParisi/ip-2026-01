package main

import "fmt"

func main() {
	fmt.Print("Digite o preço do seu carro: ")
	var preçoDoCarro float64
	fmt.Scan(&preçoDoCarro)

	fmt.Println("Para calcularmos o preço final do seu carro, diga \"sim\", ou \"não\" para as seguintes perguntas:")

	fmt.Print("\nDeseja Ar condicionado no seu veículo? ")
	var arCondicionado string
	fmt.Scan(&arCondicionado)
	if arCondicionado == "sim" {
		preçoDoCarro += 1750
		fmt.Println("O preço do seu carro com ar condicionado é: R$", preçoDoCarro)
	} else if arCondicionado == "não" {
		fmt.Println("O preço do seu carro sem ar condicionado é: R$", preçoDoCarro)
	}

	fmt.Print("\nDeseja Pintura metálica no seu veículo? ")
	var pinturaMetalica string
	fmt.Scan(&pinturaMetalica)
	if pinturaMetalica == "sim" {
		preçoDoCarro += 800
		fmt.Println("O preço do seu carro com pintura metálica é: R$", preçoDoCarro)
	} else if pinturaMetalica == "não" {
		fmt.Println("O preço do seu carro sem pintura metálica é: R$", preçoDoCarro)
	}

	fmt.Print("\nDeseja Vidro elétrico no seu veículo? ")
	var vidroEletrico string
	fmt.Scan(&vidroEletrico)
	if vidroEletrico == "sim" {
		preçoDoCarro += 1200
		fmt.Println("O preço do seu carro com vidro elétrico é: R$", preçoDoCarro)
	} else if vidroEletrico == "não" {
		fmt.Println("O preço do seu carro sem vidro elétrico é: R$", preçoDoCarro)
	}

	fmt.Print("\nDeseja Direção hidráulica no seu veículo? ")
	var direcaoHidraulica string
	fmt.Scan(&direcaoHidraulica)
	if direcaoHidraulica == "sim" {
		preçoDoCarro += 2000
		fmt.Println("O preço do seu carro com direção hidráulica é: R$", preçoDoCarro)
	} else if direcaoHidraulica == "não" {
		fmt.Println("O preço do seu carro sem direção hidráulica é: R$", preçoDoCarro)
	}

	fmt.Println("\n\nDescrição do carro: ")
	fmt.Println("Ar condicionado? ", arCondicionado)
	fmt.Println("Pintura metálica? ", pinturaMetalica)
	fmt.Println("Vidro elétrico? ", vidroEletrico)
	fmt.Println("Direção hidráulica? ", direcaoHidraulica)
	fmt.Println("Preço final do carro: R$", preçoDoCarro)

}
