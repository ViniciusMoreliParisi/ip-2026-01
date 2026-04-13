package main

import "fmt"

func main() {
	fmt.Println("Vamos!")
	fmt.Print("Digite um número X : ")
	var x float64
	fmt.Scan(&x)

	var somatorio float64
	var somafatorialimpar float64
	var fatorialimpar float64 = 1
	var somafatorialpar float64
	var fatorialpar float64 = 1

	//PROCESSO DE CÁLCULO DO SOMAFATORIALÍMPAR

	for i := 1.00; i < 20.00; i += 2.00 {
		fatorialimpar = 1
        if i <= 1 {
			fatorialimpar = 1
		} else if i > 1 {

			for c := i; c > 1; c-- {
				fatorialimpar *= c 

			}
		}
		somafatorialimpar += x / fatorialimpar
	}
	//PROCESSO DE CÁLCULO DO SOMAFATORIALPAR

	for i := 2.00; i < 20.00; i += 2.00 {
		fatorialpar = 1
        if i <= 1 {
			fatorialpar = 1
		} else if i > 1 {
			for c := i; c > 1; c-- {
				fatorialpar *= c 

			}
		}
		somafatorialpar += x / fatorialpar
	}
	somatorio = x - somafatorialimpar + somafatorialpar
	fmt.Printf("O resultado do somatório é:%.4f ", somatorio)

}
