package problemas

import "fmt"

// Lê o valor de 3 notas e calcula a média entre elas
// Sendo que a primeira nota vale 2, a segunda 3 e a terceira 5
func CalculoMedia2() {
	var notaA, notaB, notaC float64

	fmt.Scanln(&notaA)
	fmt.Scanln(&notaB)
	fmt.Scanln(&notaC)

	media := ((notaA * 2) + (notaB * 3) + (notaC * 5)) / 10

	fmt.Printf("MEDIA = %.1f\n", media)
}
