package main

import "fmt"

func main() {

	fmt.Println("Digite (horas, minutos e segundos), separados por espaço, e converterei tudo em segundos")
	var h, m, s int
	fmt.Scan(&h, &m, &s)

	total := 3600*h + 60*m + s

	fmt.Print("O total, em segundos é: ", total, "s")

}
