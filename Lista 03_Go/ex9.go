package main

import "fmt"

func main() {
	fmt.Println("Vamos trabalhar com as notas de N alunos.")
	fmt.Print("Determine N: ")
	var N int
	fmt.Scan(&N)

	var contreprovados int
	var contexame int
	var contaprovados int
	var conjmedias float64

	for cont := 1; cont <= N; cont++ {
		fmt.Println("Digite a 1° nota do aluno ", cont, ": ")
		var n1 float64
		fmt.Scan(&n1)

		fmt.Println("Digite a 2° nota do aluno ", cont, ": ")
		var n2 float64
		fmt.Scan(&n2)

		media := (n1 + n2) / 2

		conjmedias += media

		if media <= 3 {
			fmt.Println("Média Aritmética  |  Mensagem\nAté 3             |  Reprovado\nEntre 3 e 7       |  Exame\nDe 7 para cima    |  Aprovado\n\n- O aluno", cont, "de media", media, "foi REPROVADO\n\n")
			contreprovados++
		}
		if media > 3 && media < 7 {
			fmt.Println("Média Aritmética  |  Mensagem\nAté 3             |  Reprovado\nEntre 3 e 7       |  Exame\nDe 7 para cima    |  Aprovado\n\n- O aluno", cont, "de media", media, "foi EXAME\n\n")
			contexame++
		}
		if media >= 7 {
			fmt.Println("Média Aritmética  |  Mensagem\nAté 3             |  Reprovado\nEntre 3 e 7       |  Exame\nDe 7 para cima    |  Aprovado\n\n- O aluno", cont, "de media", media, "foi APROVADO\n\n")
			contaprovados++
		}
		fmt.Println("O número de alunos reprovados é: ", contreprovados)
		fmt.Println("O número de alunos em exame é: ", contexame)
		fmt.Println("O número de alunos aprovados é: ", contaprovados)
		fmt.Println("A média da classe é: ", conjmedias/float64(N))

	}
}
