package problemas

import "fmt"

func TempodeJogo(horaInicio, horaFinal int) (duracao string) {

	duracaoInt := 0
	if horaInicio == horaFinal {
		duracaoInt = 24
	} else if horaInicio > horaFinal {
		horaFinal += 24
		duracaoInt = horaFinal - horaInicio
	} else if horaFinal > horaInicio {
		duracaoInt = horaFinal - horaInicio
	}

	return fmt.Sprintf("O JOGO DUROU %d HORA(S)", duracaoInt)

}
