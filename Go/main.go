package main

import (
	"fmt"
	"problemasURI/problemas"
)

func main() {
	var numero int64
	fmt.Scanf("%d", &numero)
	fmt.Println(problemas.Mes(numero))
}
