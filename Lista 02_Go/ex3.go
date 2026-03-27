package main

import "fmt"

func main() {

	fmt.Println("Digite dois números e te darei a soma deles. Resultado > 20 = Resultado + 8, Resultado <= 20 = Resultado - 5.")

	var a, b int
	fmt.Scan(&a, &b)
	soma := a + b

	if soma > 20 {
		fmt.Print("O resultado é ", soma+8)
	} else {
		fmt.Print(" O resultado é ", soma-5)
	}

}
