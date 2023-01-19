package problemas

import "fmt"

func TesteDeSelecao() {

	var valorA, valorB, valorC, valorD int64

	fmt.Scanf("%d %d %d %d\n", &valorA, &valorB, &valorC, &valorD)

	if (valorB > valorC) && (valorD > valorA) && ((valorC + valorD) > (valorA + valorB)) && (valorA%2 == 0) && (valorD > 0) && (valorC > 0) {
		fmt.Println("Valores aceitos")
	} else {
		fmt.Println("Valores nao aceitos")
	}

}
