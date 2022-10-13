package problemas

import "fmt"

/*
	Recebe o valor de acordo com as variaveis
	Mostra 
*/
func Salario() {
	var numeroFuncionario, horasTrabalhadas, valorHora float64

	fmt.Scanln(&numeroFuncionario)
	fmt.Scanln(&horasTrabalhadas)
	fmt.Scanln(&valorHora)

	fmt.Printf("NUMBER = %.0f\nSALARY = U$ %.2f\n", numeroFuncionario, (valorHora * horasTrabalhadas))
}
