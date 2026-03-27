package main

import "fmt"

func main() {

	fmt.Print("Digite dois números, e te direi se o primeiro é divisível pelo outro: ")
	var a int
	var b int
	fmt.Scan(&a, &b)

	if a%b == 0 {
		fmt.Println("O primeiro número é divisível pelo segundo.")
	} else {
		fmt.Println("O primeiro número não é divisível pelo segundo.")
	}

}
