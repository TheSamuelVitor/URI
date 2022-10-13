package problemas

import "fmt"

func CalculaArea() {
	var valorA, valorB, valorC float64

	fmt.Scanf("%f %f %f", &valorA, &valorB, &valorC)

	fmt.Printf("TRIANGULO: %.3f\n", (valorA*valorC)/2)
	fmt.Printf("CIRCULO: %.3f\n", (valorC*valorC)*3.14159)
	fmt.Printf("TRAPEZIO: %.3f\n", ((valorA+valorB)*valorC)/2)
	fmt.Printf("QUADRADO: %.3f\n", valorB*valorB)
	fmt.Printf("RETANGULO: %.3f\n", valorA*valorB)
}
