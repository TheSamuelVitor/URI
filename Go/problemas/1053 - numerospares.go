package problemas

import "fmt"

func NumerosPares() (listaDeNumeros string) {

	for i := 4; i <= 100; i++ {
		if i == 4 {
			listaDeNumeros = fmt.Sprintf("%d", 2)
		}

		if i % 2 == 0 {
			listaDeNumeros = fmt.Sprintf("%s\n%d", listaDeNumeros, i)
		}
	}

	return

}
