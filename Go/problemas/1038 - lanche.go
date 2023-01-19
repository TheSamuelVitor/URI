package problemas

func Lanche(codigo, quantidade float64) float64 {

	var valor float64
	switch codigo {
	case 1:
		valor = 4
	case 2:
		valor = 4.5
	case 3:
		valor = 5
	case 4:
		valor = 2
	case 5:
		valor = 1.5
	}
	return valor*quantidade
}
