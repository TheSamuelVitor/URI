package problemas

import (
	"fmt"
	"math"
)

func FormuladeBhaskara(valorA, valorB, valorC float64) string {
	delta := math.Pow(valorB, 2) - 4*valorA*valorC
	if valorA == 0 || delta < 0 {
		return "Impossivel calcular\n"
	} else {
		x1 := (-valorB + math.Sqrt(delta)) / (2 * valorA)
		x2 := (-valorB - math.Sqrt(delta)) / (2 * valorA)
		return fmt.Sprintf("R1 = %.5f\nR2 = %.5f\n", x1, x2)
	}
}
