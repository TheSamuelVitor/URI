package main

import (
	"fmt"
	"problemasURI/problemas"
)

func main() {
	var horaInicial, minutoInicio, minutoFinal, horaFinal int
	fmt.Scanf("%d %d %d %d\n", &horaInicial, &minutoInicio, &horaFinal, &minutoFinal)
	fmt.Println(problemas.TempodeJogoComMinutos(horaInicial, minutoInicio, horaFinal, minutoFinal))
}
