package main

import (
	"fmt"
	"problemasURI/problemas"
)

func main() {
	var horaInicial, horaFinal int
	fmt.Scanf("%d %d\n", &horaInicial, &horaFinal)
	fmt.Println(problemas.TempodeJogo(horaInicial, horaFinal))
}
