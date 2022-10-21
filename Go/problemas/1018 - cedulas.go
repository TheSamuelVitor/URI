package problemas

import "fmt"

func Cedulas() {

	var valor int

	fmt.Scanln(&valor)

	fmt.Println(valor)

	qtd100 := valor / 100
	valor = valor - (qtd100 * 100)
	qtd50 := valor / 50
	valor = valor - (qtd50 * 50)
	qtd20 := valor / 20
	valor = valor - (qtd20 * 20)
	qtd10 := valor / 10
	valor = valor - (qtd10 * 10)
	qtd5 := valor / 5
	valor = valor - (qtd5 * 5)
	qtd2 := valor / 2
	valor = valor - (qtd2 * 2)

	fmt.Printf("%d nota(s) de R$ 100,00\n", qtd100)
	fmt.Printf("%d nota(s) de R$ 50,00\n", qtd50)
	fmt.Printf("%d nota(s) de R$ 20,00\n", qtd20)
	fmt.Printf("%d nota(s) de R$ 10,00\n", qtd10)
	fmt.Printf("%d nota(s) de R$ 5,00\n", qtd5)
	fmt.Printf("%d nota(s) de R$ 2,00\n", qtd2)
	fmt.Printf("%d nota(s) de R$ 1,00\n", valor)

}