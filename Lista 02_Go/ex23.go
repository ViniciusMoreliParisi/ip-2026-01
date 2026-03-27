package main

import "fmt"

func main() {

	fmt.Print("Cidadão, digite sua idade: ")
	var idade int
	fmt.Scan(&idade)

	var classe string

	switch {
	case idade < 16 && idade >= 0:
		classe = "Não-eleitor"

	case idade > 18 && idade < 65:
		classe = "Eleitor obrigatório"

	case idade > 16 && idade < 18 || idade > 65:
		classe = "Eleitor facultativo"

	default:
		fmt.Println("Insira uma idade válida!")
		return
	}

	fmt.Println("\nCidadão, sua classe eleitoral é: ", classe)

}
