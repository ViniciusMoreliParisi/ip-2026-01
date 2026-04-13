package main

import (
	"fmt"
)

func main() {
	fmt.Println("Declare o peso dos bois, e dar-te-ei o peso do boi mais pesado, do mais leve.")

	var id_maior int
	var maior float64 = 0
	var id_menor int
	var menor float64 = 10000
	for i := 1; i <= 90; i++ {

		fmt.Printf("Digite id do %d° boi: ", i)
		var id int
		fmt.Scan(&id)
		fmt.Printf("Digite peso do %d° boi: ", i)
		var n float64
		fmt.Scan(&n)

		if n > maior {
			maior = n
			id_maior = id
		}
		if n < menor {
			menor = n
			id_menor = id
		}

	}
	fmt.Printf("O boi %d é o mais pesado e pesa %.2f kg.\n", id_maior, maior)
	fmt.Printf("O boi %d é o mais leve e pesa %.2f kg.\n", id_menor, menor)

}
