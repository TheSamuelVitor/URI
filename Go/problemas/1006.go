package problemas

import "fmt"

func CalculoMedia2() {
	var notaA, notaB, notaC float64

	fmt.Scanln(&notaA)
	fmt.Scanln(&notaB)
	fmt.Scanln(&notaC)

	media := ((notaA*2) + (notaB*3) + (notaC*5) )/10

	fmt.Printf("MEDIA = %.1f\n", media)
}