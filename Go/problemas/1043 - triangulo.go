package problemas

import "fmt"

func Triangulo(ladoA, ladoB, ladoC float64) (resultado string) {

	if ladoA < (ladoB+ladoC) && ladoB < (ladoC+ladoA) && ladoC < (ladoA+ladoB) {
		perimetro := ladoA + ladoB + ladoC
		resultado = fmt.Sprintf("Perímetro = %.1f\n", perimetro)
		return
	} else {
		area := ((ladoA + ladoB) * ladoC) / 2
		resultado = fmt.Sprintf("Area = %.1f\n", area)
		return
	}

}
