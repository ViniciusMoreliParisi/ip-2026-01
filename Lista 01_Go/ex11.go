package main

import "fmt"

func main() {
	fmt.Println("Vamos ver se seu número é divisível por 3 e 5?")
	fmt.Print("Digite o número: ")
	var n int
	fmt.Scan(&n)

	if n%3 == 0 && n%5 == 0 {
		fmt.Println("O número é divisível sim! :)")
	} else {
		fmt.Println("O número não é divisível! :(")
	}

}
