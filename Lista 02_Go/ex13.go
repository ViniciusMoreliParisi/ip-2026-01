package main

import "fmt"

func main() {

	fmt.Print("Digite um número de três algarismos, e te darei o valor da casa decimal. \nLembre-se de dar espaço a cada dígito*\n")
	var c, d, u int
	fmt.Scan(&c, &d, &u)
	fmt.Println("A casa decimal do número apresentado é: ", d)
}
