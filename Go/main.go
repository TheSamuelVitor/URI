package main

import (
	"fmt"
	"problemasURI/problemas"
)

func main() {
	var valor1, valor2 int
	fmt.Scanf("%d\n%d", &valor1, &valor2)
	fmt.Println(problemas.SomaImpares(valor1,valor2))
}
