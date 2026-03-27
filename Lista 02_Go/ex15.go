package main

import "fmt"

func main() {

	fmt.Print("Digite uma data, e a reescreverei com o mês por extenso :) \nLembre-se de digitar o dia, mês e ano, separados por espaços: ")
	var dia, mes, ano int
	fmt.Scan(&dia, &mes, &ano)

	var mesNome string
	if mes == 1 {
		mesNome = "janeiro"
	} else if mes == 2 {
		mesNome = "fevereiro"
	} else if mes == 3 {
		mesNome = "março"
	} else if mes == 4 {
		mesNome = "abril"
	} else if mes == 5 {
		mesNome = "maio"
	} else if mes == 6 {
		mesNome = "junho"
	} else if mes == 7 {
		mesNome = "julho"
	} else if mes == 8 {
		mesNome = "agosto"
	} else if mes == 9 {
		mesNome = "setembro"
	} else if mes == 10 {
		mesNome = "outubro"
	} else if mes == 11 {
		mesNome = "novembro"
	} else if mes == 12 {
		mesNome = "dezembro"
	}

	fmt.Print("A data é: ", dia, " de ", mesNome, " de ", ano)
}
