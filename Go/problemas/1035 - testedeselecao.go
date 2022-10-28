package problemas

func TesteDeSelecao(valorA, valorB, valorC, valorD int) string {
	if (valorB > valorC) && (valorD > valorA) && ((valorC+valorD)>(valorA+valorB)) && (valorA % 2 == 0) && (valorD > 0) && (valorC > 0) {
		return "Valores aceitos"
	}
	
	return "Valores nao aceitos"
}