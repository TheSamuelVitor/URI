package problemas

import "fmt"

func GastodeCombustivel() {

	var tempoDeViagem, velocidadeMedia float64

	fmt.Scanln(&tempoDeViagem)
	fmt.Scanln(&velocidadeMedia)

	quantidadeDeLitros := (velocidadeMedia/12) * tempoDeViagem

	fmt.Printf("%.3f\n", quantidadeDeLitros)
}