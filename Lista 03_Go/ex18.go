package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Vamos dispor a diagonal principal de uma matriz 10X10?")
	fmt.Print("Digite sim, ou não: ")
	var sim string
	fmt.Scan(&sim)

	sim = strings.ToUpper(sim)

	if sim != "SIM" {
		fmt.Println("Ok, até a próxima!")
	} else {

		var coluna int = 1

		for linha := 1; linha <= 10; linha++ {
			for coluna = 1; coluna <= 10; coluna++ {
				if coluna == linha {
					fmt.Printf("(%d, %d) ", linha, coluna)
				} else {
					fmt.Print("    ")
				}
				if coluna == 10 {
					fmt.Print("\n")
				}
			}
		}
	}
}
