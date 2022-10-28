package problemas

import (
	"fmt"
	"math"
)

func NotasEMoedas() {
	var dinheiro float64
	var notas100, notas50, notas20, notas10, notas5, notas2 float64
	var moedas100, moedas50, moedas25, moedas10, moedas5, moedas1 float64

	fmt.Scanln(&dinheiro)

	notas100 = math.Floor(dinheiro / 100)
	notas50 = math.Floor((dinheiro - (notas100 * 100)) / 50)
	notas20 = math.Floor((dinheiro - (notas100 * 100) - (notas50 * 50)) / 20)
	notas10 = math.Floor((dinheiro - (notas100 * 100) - (notas50 * 50) - (notas20 * 20)) / 10)
	notas5 = math.Floor((dinheiro - (notas100 * 100) - (notas50 * 50) - (notas20 * 20) - (notas10 * 10)) / 5)
	notas2 = math.Floor((dinheiro - (notas100 * 100) - (notas50 * 50) - (notas20 * 20) - (notas10 * 10) - (notas5 * 5)) / 2)

	dinheiro2 := math.Floor((dinheiro - (notas100 * 100) - (notas50 * 50) - (notas20 * 20) - (notas10 * 10) - (notas5 * 5) - (notas2 * 2)) * 100)

	moedas100 = math.Floor(dinheiro2 / 100)
	moedas50 = math.Floor((dinheiro2 - (moedas100 * 100)) / 50)
	moedas25 = math.Floor((dinheiro2 - (moedas100 * 100) - (moedas50 * 50)) / 25)
	moedas10 = math.Floor((dinheiro2 - (moedas100 * 100) - (moedas50 * 50) - (moedas25 * 20)) / 10)
	moedas5 = math.Floor((dinheiro2 - (moedas100 * 100) - (moedas50 * 50) - (moedas25 * 25) - (moedas10 * 10)) / 5)
	moedas1 = math.Floor(dinheiro2 - (moedas100 * 100) - (moedas50 * 50) - (moedas25 * 25) - (moedas10 * 10) - (moedas5 * 5))

	fmt.Printf("NOTAS:\n")
	fmt.Printf("%.0f nota(s) de R$ 100.00\n", notas100)
	fmt.Printf("%.0f nota(s) de R$ 50.00\n", notas50)
	fmt.Printf("%.0f nota(s) de R$ 20.00\n", notas20)
	fmt.Printf("%.0f nota(s) de R$ 10.00\n", notas10)
	fmt.Printf("%.0f nota(s) de R$ 5.00\n", notas5)
	fmt.Printf("%.0f nota(s) de R$ 2.00\n", notas2)

	fmt.Printf("MOEDAS:\n")
	fmt.Printf("%.0f moeda(s) de R$ 1.00\n", moedas100)
	fmt.Printf("%.0f moeda(s) de R$ 0.50\n", moedas50)
	fmt.Printf("%.0f moeda(s) de R$ 0.25\n", moedas25)
	fmt.Printf("%.0f moeda(s) de R$ 0.10\n", moedas10)
	fmt.Printf("%.0f moeda(s) de R$ 0.05\n", moedas5)
	fmt.Printf("%.0f moeda(s) de R$ 0.01\n", moedas1)
}
