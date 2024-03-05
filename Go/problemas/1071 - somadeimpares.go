package problemas

// Retorna a soma dos ímpares consecutivos entre os valores dados
func SomaImpares(valor1, valor2 int) (soma int) {

	if valor2%2 != 0 {
		valor2++
	}

	for i := valor2; i < valor1; i++ {
		if i%2 != 0 {
			soma += i
		}
	}

	return
}
