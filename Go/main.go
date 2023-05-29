package main

import (
	"fmt"
	"problemasURI/problemas"
)

func main() {
	var salario float64
	fmt.Scanf("%f", &salario)
	problemas.AumentoDeSalario(salario)
}
