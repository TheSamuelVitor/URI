package problemas

import "math"

func TiposDeTriangulo(ladoA, ladoB, ladoC float64) (resultado string) {
	
	var numero1, numero2, numero3 float64

	if ladoA > ladoB && ladoA > ladoC {
		numero1 = ladoA
		if ladoB > ladoC {
			numero2 = ladoB
			numero3 = ladoC
		} else {
			numero2 = ladoC
			numero3 = ladoB
		}
	} else if ladoB > ladoA && ladoB > ladoC {
		numero1 = ladoB
		if ladoA > ladoC {
			numero2 = ladoA
			numero3 = ladoC
		} else {
			numero2 = ladoC
			numero3 = ladoA
		}
	} else if ladoC > ladoB && ladoC > ladoA {
		numero1 = ladoC
		if ladoA > ladoB {
			numero2 = ladoA
			numero3 = ladoB
		} else {
			numero2 = ladoB
			numero3 = ladoA
		}
	}

	if numero1 >= numero2+numero3 {
		resultado = "NAO FORMA TRIANGULO"
	} else {
		if math.Pow( numero1,2) == math.Pow(numero2,2)+math.Pow(numero3,2) {
			resultado = "TRIANGULO RETANGULO\n"
		} else if math.Pow( numero1,2) > math.Pow(numero2,2)+math.Pow(numero3,2) {
			resultado = "TRIANGULO OBTUSANGULO\n"	
		} else if math.Pow( numero1,2) < math.Pow(numero2,2)+math.Pow(numero3,2) {
			resultado = "TRIANGULO ACUTANGULO\n"
		}

		if ladoA == ladoB && ladoB == ladoC {
			resultado += "TRIANGULO EQUILATERO"
		} else if ladoA == ladoC || ladoA == ladoB || ladoB == ladoC {
			resultado += "TRIANGULO ISOSCELES
		}
	}
	return

}
