package main

import (
	"fmt"
)

func main() {
	var valorA, valorB int
	fmt.Scanf("%d %d\n", &valorA, &valorB)
	fmt.Print(Multiplos(valorA, valorB))
}

func Multiplos(valorA, valorB int) string {

	if valorA%valorB == 0 || valorB%valorA == 0 {
		return "Sao Multiplos\n"
	} else {
		return "Nao sao Multiplos\n"
	}

}
