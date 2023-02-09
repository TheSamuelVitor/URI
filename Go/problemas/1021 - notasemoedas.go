package problemas

import (
	"fmt"
	"math"
)

func NotasEMoedas(dinheiro float64) string {
	var notas100, notas50, notas20, notas10, notas5, notas2 float64
	var moedas100, moedas50, moedas25, moedas10, moedas5, moedas1 float64

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
	moedas10 = math.Floor((dinheiro2 - (moedas100 * 100) - (moedas50 * 50) - (moedas25 * 25)) / 10)
	moedas5 = math.Floor((dinheiro2 - (moedas100 * 100) - (moedas50 * 50) - (moedas25 * 25) - (moedas10 * 10)) / 5)
	moedas1 = math.Floor(dinheiro2 - (moedas100 * 100) - (moedas50 * 50) - (moedas25 * 25) - (moedas10 * 10) - (moedas5 * 5))

	return fmt.Sprintf("NOTAS:\n%.0f nota(s) de R$ 100.00\n%.0f nota(s) de R$ 50.00\n%.0f nota(s) de R$ 20.00\n%.0f nota(s) de R$ 10.00\n%.0f nota(s) de R$ 5.00\n%.0f nota(s) de R$ 2.00\nMOEDAS:\n%.0f moeda(s) de R$ 1.00\n%.0f moeda(s) de R$ 0.50\n%.0f moeda(s) de R$ 0.25\n%.0f moeda(s) de R$ 0.10\n%.0f moeda(s) de R$ 0.05\n%.0f moeda(s) de R$ 0.01\n", notas100, notas50, notas20, notas10, notas5, notas2, moedas100, moedas50, moedas25, moedas10, moedas5, moedas1)
}
