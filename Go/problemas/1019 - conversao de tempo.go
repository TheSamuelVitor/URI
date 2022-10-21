package problemas

import "fmt"

func ConversaoDeTempo() {
	var segundos int

	fmt.Scanln(&segundos)

	horas := segundos / 3600
	segundos = segundos - (3600 * horas)

	minutos := segundos / 60
	segundos = segundos - (minutos * 60)

	fmt.Printf("%d:%d:%d\n", horas, minutos, segundos)
}