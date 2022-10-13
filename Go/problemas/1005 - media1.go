package problemas

import "fmt"

// Lê dois valores da linha de comando e mostra uma media entre eles
// sendo que o primeiro valor possui peso de 3.5 e o segundo de 7.5
func CalculoMedia() {
	var valor1, valor2 float64

	fmt.Scanln(&valor1)
	fmt.Scanln(&valor2)

	media := ((valor1*3.5) + (valor2*7.5))/11

	fmt.Printf("MEDIA = %.5f\n", media)
}