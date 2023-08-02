package problemas

// funcao que recebe o valor de salario e retorna o salario com a modificacao
func AumentoDeSalario(salario float64) (reajuste float64, percentual string) {
	if salario >= 0 && salario <= 400 {
		reajuste = salario * 0.15
		percentual = "15%"
	} else if salario >= 400.01 && salario <= 800 {
		reajuste = salario * 0.12
		percentual = "12%"
	} else if salario >= 800.01 && salario <= 1200 {
		reajuste = salario * 0.1
		percentual = "10%"
	} else if salario >= 1200.01 && salario <= 2000 {
		reajuste = salario * 0.07
		percentual = "7%"
	} else if salario > 2000 {
		reajuste = salario * 0.04
		percentual = "4%"
	}

	return reajuste, percentual
}
