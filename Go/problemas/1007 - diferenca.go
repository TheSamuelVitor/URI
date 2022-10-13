package problemas

import "fmt"

// Lê quatro valores e mostra a diferenca entre eles
func Diferenca() {
	var valorA, valorB, valorC, valorD int

	fmt.Scanln(&valorA)
	fmt.Scanln(&valorB)
	fmt.Scanln(&valorC)
	fmt.Scanln(&valorD)

	diferenca := (valorA*valorB)-(valorC*valorD)

	fmt.Println("DIFERENCA =", diferenca)
}