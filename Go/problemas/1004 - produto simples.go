package problemas

import "fmt"

// Recebe dois valores e mostra a multiplicação entre eles
func CalculaProduto() {
	var primeiroNumero, segundoNumero int

	fmt.Scanln(&primeiroNumero)
	fmt.Scanln(&segundoNumero)

	fmt.Println("PROD =", primeiroNumero*segundoNumero)
}
