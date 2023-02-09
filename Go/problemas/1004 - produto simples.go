package problemas

import "fmt"

// Recebe dois valores e mostra a multiplicação entre eles
func CalculaProduto(primeiroNumero, segundoNumero int) string {
	return fmt.Sprintf("PROD = %v", primeiroNumero*segundoNumero)
}
