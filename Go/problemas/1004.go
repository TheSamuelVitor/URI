package problemas

import "fmt"

func CalculaProduto() {
	var primeiroNumero, segundoNumero int

	fmt.Scanln(&primeiroNumero)
	fmt.Scanln(&segundoNumero)

	fmt.Println("PROD =", primeiroNumero*segundoNumero)
}
