package problemas

func Multiplos(valorA, valorB int) string {

	if valorA % valorB == 0 || valorB % valorA == 0 {
		return "Sao Multiplos\n"
	} else {
		return "Nao sao Multiplos\n"
	}

}
