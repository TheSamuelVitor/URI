package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e, f float64
	fmt.Scanf("%f\n%f\n%f\n%f\n%f\n%f", &a, &b, &c, &d, &e, &f)
	fmt.Println(NumerosPositivos([]float64{a, b, c, d, e, f}))
}

func NumerosPositivos(numeros []float64) string {

	var numerosPositivos int64

	for _, numero := range numeros {
		if numero > 0 {
			numerosPositivos += 1
		}
	}

	return fmt.Sprintf("%d valores positivos", numerosPositivos)
}
