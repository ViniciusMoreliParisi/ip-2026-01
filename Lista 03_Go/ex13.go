package main

import "fmt"

func main() {
	fmt.Println("Vamos determinar o valor de H?")

	var num int = 1
	var den int = 1
	var h int

	fmt.Println("H = ")

	for i := 50; i > 0; i-- {

		h += num / den

		fmt.Print(num, "/", den)
		if i > 1 {
			fmt.Println(" + ")
		}
		num += 2
		den += 1
	}

	fmt.Printf("\nH = %d", h)

}
