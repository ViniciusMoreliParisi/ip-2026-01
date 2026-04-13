package main

import "fmt"

func main() {
	fmt.Println("Vamos formar um P.A?")
	fmt.Print("Digite a quantidade de termos: ")
	var n int
	fmt.Scan(&n)

	var a int = 0
	var b int = 1

	for i := n; i > 0; i-- {
		a += b
		fmt.Print(a)
		if i > 1 {
			fmt.Print(", ")
		}
		b += 2
	}

}
