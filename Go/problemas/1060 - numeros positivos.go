package problemas

import "fmt"

func NumerosPositivos(numeros []int64) string {

	var numerosPositivos int64

	for _, numero := range numeros {
		if numero > 0 {
			numerosPositivos += 1
		}
	}

	return fmt.Sprintf("%d valores positivos", numerosPositivos)
}
