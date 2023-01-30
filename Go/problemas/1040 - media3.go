package problemas

import "fmt"

func Media3(n1, n2, n3, n4 float64) string {

	var media, nExame float64
	media = ((n1 * 2) + (n2 * 3) + (n3 * 4) + (n4)) / 10

	fmt.Printf("Media: %.1f\n", media)
	if media >= 7 {
		return "Aluno aprovado.\n"
	} else if media < 5 {
		return "Aluno reprovado.\n"
	} else if media >= 5 && media <= 6.9 {
		fmt.Printf("Aluno em exame.\n")
		fmt.Scanf("%f", &nExame)
		fmt.Printf("Nota do exame: %.1f\n", nExame)
		media = (media + nExame) / 2
		if media >= 5 {
			fmt.Printf("Aluno aprovado.\n")
		} else {
			fmt.Printf("Aluno reprovado.\n")
		}
		return fmt.Sprintf("Media final: %.1f\n", media)
	}

	return ""

}
