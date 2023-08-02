package main

import (
	"fmt"
	"problemasURI/problemas"
)

func main() {
	var salario float64
	fmt.Scanf("%f", &salario)
	var reajuste, porcentagem = problemas.AumentoDeSalario(salario)
	fmt.Printf("Novo salario: %.2f\n", salario+reajuste)
	fmt.Printf("Reajuste ganho: %.2f\n", reajuste)
	fmt.Printf("Em percentual: %s\n", porcentagem)
}
