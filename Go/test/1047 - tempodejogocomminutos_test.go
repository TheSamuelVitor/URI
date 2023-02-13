package test

import (
	"fmt"
	"problemasURI/assert"
	"problemasURI/problemas"
	"testing"
)

func TestTempoComMinutos(t *testing.T) {
	tests := []struct {
		got      string
		expected string
	}{
		// casos de testes para hora final igual a hora inicial
		{
			problemas.TempodeJogoComMinutos(7, 7, 7, 7),
			"O JOGO DUROU 24 HORA(S) E 0 MINUTO(S)",
		},
		{
			problemas.TempodeJogoComMinutos(7, 7, 7, 8),
			"O JOGO DUROU 0 HORA(S) E 1 MINUTO(S)",
		},
		{
			problemas.TempodeJogoComMinutos(7, 8, 7, 7),
			"O JOGO DUROU 23 HORA(S) E 59 MINUTO(S)",
		}, 

		// casos de testes para hora final é maior a hora inicial
		{
			problemas.TempodeJogoComMinutos(7, 8, 9, 7),
			"O JOGO DUROU 1 HORA(S) E 59 MINUTO(S)",
		},
		{
			problemas.TempodeJogoComMinutos(7, 8, 9, 8),
			"O JOGO DUROU 2 HORA(S) E 0 MINUTO(S)",
		},
		{
			problemas.TempodeJogoComMinutos(7, 8, 9, 10),
			"O JOGO DUROU 2 HORA(S) E 2 MINUTO(S)",
		},

		// casos de testes para hora final é menor a hora inicial
		{
			problemas.TempodeJogoComMinutos(9, 8, 7, 8),
			"O JOGO DUROU 22 HORA(S) E 0 MINUTO(S)",
		},
		{
			problemas.TempodeJogoComMinutos(9, 8, 7, 10),
			"O JOGO DUROU 22 HORA(S) E 2 MINUTO(S)",
		},
		{
			problemas.TempodeJogoComMinutos(9, 8, 7, 7),
			"O JOGO DUROU 21 HORA(S) E 59 MINUTO(S)",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Teste %d", i+1), func(t *testing.T) {
			assert.Equal(t, test.got, test.expected)
		})
	}

}
