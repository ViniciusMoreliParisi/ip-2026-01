package main

import (
	"fmt"
)

func main() {
	fmt.Println("Vamos transformar um número base 8 em um número de base 10!")
	fmt.Print("Digite o nosso número que será formatado: ")
	var n int
	fmt.Scan(&n)
   
   if n >= 10 {
   var resto int = n % 10 
   var quociente int = (n - resto) / 10 
   n = n - (2 * quociente)
   }
	fmt.Printf("O número %.d em decimal é: %d ", n, n)

}
