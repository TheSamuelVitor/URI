package problemas

import "fmt"

func Esfera() {
	var raio float64

	fmt.Scanln(&raio)

	pi := 3.14159

	esfera := (4.0/3) * pi * (raio * raio * raio) 
	fmt.Printf("VOLUME = %.3f\n", esfera)
}
