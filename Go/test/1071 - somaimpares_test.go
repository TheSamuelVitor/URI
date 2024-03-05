package test

import (
	"problemasURI/assert"
	"problemasURI/problemas"
	"testing"
)

func TestSomaImpares(t *testing.T) {
	casosPossiveis := []struct {
		nome          string
		valor1        int
		valor2        int
		valorEsperado int
	}{
		{
			"Teste5",
			6,
			-5,
			5,
		},
		{
			"Teste13",
			15,
			12,
			13,
		},
		{
			"Teste0",
			12,
			12,
			0,
		},
	}

	for _, caso := range casosPossiveis {
		t.Run(caso.nome, func(t *testing.T) {
			recebido := problemas.SomaImpares(caso.valor1, caso.valor2)
			assert.Equal(t, recebido,caso.valorEsperado)
		})
	}
}
