package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Vamos dispor a sombra inferior da diagonal principal de uma matriz 10X10?")
	fmt.Print("Digite sim, ou não: ")
	var sim string
	fmt.Scan(&sim)

	sim = strings.ToUpper(sim)

	if sim != "SIM" {
		fmt.Println("Ok, até a próxima!")
	} else {

		cond := 1

		for linha := 1; linha <= 10; linha++ {
			for coluna := 1; coluna <= 10; coluna++ {
				if linha <= cond && coluna <= cond {
					fmt.Printf("(%d, %d) ", linha, coluna)
				} else {
					fmt.Print("    ")
					if coluna == 10 {
						fmt.Print("\n")
					}
				}
			}
			cond++
		}
	}
}
