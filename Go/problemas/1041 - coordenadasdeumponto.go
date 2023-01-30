package problemas

func CoordenadasDeUmPonto(coordenadaX, coordenadaY float64) string {

	if (coordenadaX > 0) && (coordenadaY > 0) {
		return "Q1"
	} else if (coordenadaX < 0) && (coordenadaY > 0) {
		return "Q2"
	} else if (coordenadaX < 0) && (coordenadaY < 0) {
		return "Q3"
	} else if (coordenadaX > 0) && (coordenadaY < 0) {
		return "Q4"
	} else if (coordenadaX == 0) && (coordenadaY == 0) {
		return "Origem"
	} else if coordenadaX == 0 {
		return "Eixo Y"
	} else if coordenadaY == 0 {
		return "Eixo X"
	}

	return ""

}
