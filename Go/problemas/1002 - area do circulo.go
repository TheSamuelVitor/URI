package problemas

import (
	"fmt"
)

// Recebe o raio do círculo e mostra a área de acordo com o valor dado
func CalculoAreadoCirculo(raio float64) string {
	area := fmt.Sprintf("%.4f", (raio*raio)*3.14159)
	return area
}
