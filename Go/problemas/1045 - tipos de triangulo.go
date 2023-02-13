package problemas

import (
	"math"
)

func TiposDeTriangulo(ladoA, ladoB, ladoC float64) (resultado string) {

	var numero1, numero2, numero3 float64

	if ladoA > ladoB && ladoA > ladoC {
		numero1 = ladoA
		numero2 = ladoC
		numero3 = ladoB
	} else if ladoB > ladoA && ladoB > ladoC {
		numero1 = ladoB
		numero2 = ladoA
		numero3 = ladoC
	} else if ladoC > ladoB && ladoC > ladoA {
		numero1 = ladoC
		numero2 = ladoA
		numero3 = ladoB
	} else {
		numero1 = ladoA
		numero2 = ladoC
		numero3 = ladoB
	}

	if numero1 >= numero2+numero3 {
		resultado = "NAO FORMA TRIANGULO"
		return
	} else {
		if math.Pow(numero1, 2) == math.Pow(numero2, 2)+math.Pow(numero3, 2) {
			resultado = "TRIANGULO RETANGULO"
		} else if math.Pow(numero1, 2) > math.Pow(numero2, 2)+math.Pow(numero3, 2) {
			resultado = "TRIANGULO OBTUSANGULO"
		} else if math.Pow(numero1, 2) < math.Pow(numero2, 2)+math.Pow(numero3, 2) {
			resultado = "TRIANGULO ACUTANGULO"
		}

		if ladoA == ladoB && ladoB == ladoC {
			resultado += "\nTRIANGULO EQUILATERO"
		} else if ladoA == ladoC || ladoA == ladoB || ladoB == ladoC {
			resultado += "\nTRIANGULO ISOSCELES"
		}
		return
	}

}
