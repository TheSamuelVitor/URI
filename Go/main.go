package main

import (
	"fmt"
	"problemasURI/problemas"
)

func main() {
	var ladoA, ladoB, ladoC float64
	fmt.Scanf("%f %f %f\n", &ladoA, &ladoB, &ladoC)
	fmt.Println(problemas.TiposDeTriangulo(ladoA, ladoB, ladoC))
}
