package problemas

import (
	"fmt"
)

// Recebe o raio do círculo e mostra a área de acordo com o valor dado
func CalculoAreadoCirculo(raio float64) string {
	// var raio float64

	// fmt.Scanln(&raio)

	// fmt.Printf("A=%.4f\n", (raio*raio)*3.14159)
	area := fmt.Sprintf("%.4f", (raio*raio)*3.14159)
	return area
}
