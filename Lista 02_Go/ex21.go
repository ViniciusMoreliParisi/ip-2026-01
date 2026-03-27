package main

import "fmt"

func main() {

	fmt.Print("Aluno, digite seu número de identificação (5 números): ")
	var id int
	fmt.Scan(&id)

	fmt.Println("Vamos, agora, compartilhar suas notas...")

	var nota1 float64
	fmt.Print("Digite sua 1° nota: ")
	fmt.Scan(&nota1)

	var nota2 float64
	fmt.Print("Digite sua 2° nota: ")
	fmt.Scan(&nota2)

	var nota3 float64
	fmt.Print("Digite sua 3° nota: ")
	fmt.Scan(&nota3)

	var mediaExercicios float64
	fmt.Print("Por fim, me diga qual foi a media dos seus exercícios: ")
	fmt.Scan(&mediaExercicios)

	mediaFinal := (nota1 + nota2*2 + nota3*3 + mediaExercicios) / 7

	var conceito string

	if mediaFinal >= 9 && mediaFinal <= 10 {
		conceito = "A"
	} else if mediaFinal >= 7.5 && mediaFinal < 9 {
		conceito = "B"
	} else if mediaFinal >= 6 && mediaFinal < 7.5 {
		conceito = "C"
	} else if mediaFinal >= 4 && mediaFinal < 6 {
		conceito = "D"
	} else if mediaFinal < 4 {
		conceito = "E"
	}

	var aprovação string
	switch conceito {
	case "A", "B", "C":
		aprovação = "APROVADO"

	case "D", "E":
		aprovação = "REPROVADO"
		break
	}

	fmt.Print("\n Aluno, de id: ", id, "\n- Nota 01: ", nota1, "\n- Nota 02: ", nota2, "\n- Nota 03: ", nota3)
	fmt.Printf("\n- Média de aproveitamento: %2.f", mediaFinal)
	fmt.Print("\n- Conceito/Status: ", conceito, " / ", aprovação)
}
