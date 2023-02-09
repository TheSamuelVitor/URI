package problemas

import "fmt"

// Lê dois valores da linha de comando e mostra uma media entre eles
// sendo que o primeiro valor possui peso de 3.5 e o segundo de 7.5
func CalculoMedia(valor1, valor2 float64) string {

	media := ((valor1 * 3.5) + (valor2 * 7.5)) / 11

	return fmt.Sprintf("MEDIA = %.5f", media)

}
