package test

import (
	"problemasURI/assert"
	"problemasURI/problemas"
	"testing"
)

func TestTempoDeJogo(t *testing.T) {
	tempoTests := []struct {
		name     string
		got      string
		expected string
	}{
		{
			"Teste 1",
			problemas.TempodeJogo(16, 2),
			"O JOGO DUROU 10 HORA(S)",
		}, {
			"Teste 2",
			problemas.TempodeJogo(0, 0),
			"O JOGO DUROU 24 HORA(S)",
		}, {
			"Teste 3",
			problemas.TempodeJogo(2, 16),
			"O JOGO DUROU 14 HORA(S)",
		},
	}

	for _, test := range tempoTests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.got, test.expected)
		})
	}
}
