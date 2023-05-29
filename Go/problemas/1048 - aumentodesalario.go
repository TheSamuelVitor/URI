package problemas

import "fmt"

// funcao que recebe o valor de salario e retorna o salario com a modificacao
func AumentoDeSalario(salario float64) (reajuste float64, percentual string) {
	if (salario >= 0 && salario <= 400) {
		reajuste = salario * 0.15
		salario = salario + reajuste
		percentual = "15%"
	} else if (salario >= 400.01 && salario <= 800) {
		reajuste = salario * 0.12
		salario = salario + reajuste
		percentual = "12%"
	}	else if (salario >= 800.01 && salario <= 1200) {
		reajuste = salario * 0.1
		salario = salario + reajuste
		percentual = "10%"
	}	else if (salario >= 1200.01 && salario <= 2000) {
		reajuste = salario * 0.07
		salario = salario + reajuste
		percentual = "7%"
	}	else if (salario > 2000) {
		reajuste = salario * 0.05
		salario = salario + reajuste
		percentual = "5%"
	}
	
	fmt.Println("Novo salario: ", salario)
	fmt.Println("Reajuste ganho: ", reajuste)
	fmt.Println("Em percentual: ", percentual)
	return salario, reajuste, percentual
}