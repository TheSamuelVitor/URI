package problemas

import "fmt"

// Lê dois valores e retonar a soma entre eles
func SomaValores() {
	var a,b int

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	fmt.Println("X =", a+b)
}