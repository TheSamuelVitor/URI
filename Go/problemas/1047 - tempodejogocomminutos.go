package problemas

import (
	"fmt"
)

func TempodeJogoComMinutos(horaInicio, minutoInicio, horaFinal, minutoFinal int) (duracao string) {

	duracaoHora := 0
	duracaoMinuto := 0
	if horaInicio == horaFinal {
		if minutoInicio == minutoFinal {
			duracaoHora = 24
			duracaoMinuto = 0
		} else if minutoFinal > minutoInicio {
			duracaoHora = 0
			duracaoMinuto = minutoFinal - minutoInicio
		} else if minutoInicio > minutoFinal {
			duracaoHora = 23
			duracaoMinuto = (minutoFinal+60) - minutoInicio
		}
	} else if horaInicio < horaFinal {
		if minutoFinal > minutoInicio {
			duracaoHora = horaFinal - horaInicio
			duracaoMinuto = minutoFinal - minutoInicio
		} else if minutoFinal < minutoInicio {
			duracaoHora = (horaFinal - horaInicio) - 1
			duracaoMinuto = (minutoFinal+60) - minutoInicio
		} else if minutoFinal == minutoInicio {
			duracaoHora = horaFinal - horaInicio
			duracaoMinuto = 0
		}
	} else if horaInicio > horaFinal {
		if minutoFinal == minutoInicio {
			duracaoHora = (horaFinal + 24) - horaInicio
			duracaoMinuto = 0
		} else if minutoFinal > minutoInicio {
			duracaoHora = (horaFinal + 24) - horaInicio
			duracaoMinuto = minutoFinal - minutoInicio
		} else if minutoFinal < minutoInicio {
			duracaoHora = ((horaFinal + 24) - horaInicio) - 1
			duracaoMinuto = (minutoFinal + 60) - minutoInicio
		}
	}

	return fmt.Sprintf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)", duracaoHora, duracaoMinuto)
}
