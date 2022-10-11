package problemas

import "fmt"

func CalculoAreadoCirculo() {
	var raio float64

	fmt.Scanln(&raio)

	fmt.Printf("A=%.4f\n", (raio*raio)*3.14159)
}