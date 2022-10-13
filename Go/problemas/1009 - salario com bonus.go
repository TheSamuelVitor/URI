package problemas

import "fmt"

func SalarioComBonus() {
	var nomeVendedor string
	var salarioFixo, vendasEfetuadas float64

	fmt.Scanln(&nomeVendedor)
	fmt.Scanln(&salarioFixo)
	fmt.Scanln(&vendasEfetuadas)

	salario := salarioFixo + (vendasEfetuadas * 15/100)

	fmt.Printf("TOTAL = R$ %.2f\n", salario)
}