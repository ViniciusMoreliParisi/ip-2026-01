package main

import "fmt"

func main() {

	fmt.Print("Digite uma idade, e a classificarei: ")
	var idade int
	fmt.Scan(&idade)

	if idade < 0 {
		fmt.Print("Digite uma idade válida (maior que zero): ")
	} else {
		if idade <= 2 {
			fmt.Println("A idade apresentada é de um Recém-nascido.")
		}
		if idade > 2 && idade <= 11 {
			fmt.Println("A idade apresentada é de uma Criança.")
		}
		if idade > 11 && idade <= 19 {
			fmt.Println("A idade apresentada é de um Adolescente.")
		}
		if idade > 19 && idade <= 55 {
			fmt.Println("A idade apresentada é de um Adulto.")
		}
		if idade > 55 {
			fmt.Println("A idade apresentada é de um Idoso.")
		}
	}
}
