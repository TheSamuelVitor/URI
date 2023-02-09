package test

import (
	"problemasURI/assert"
	"problemasURI/problemas"
	"testing"
)

func TestNotasEMoedas(t *testing.T) {

	notasemoedastests := []struct {
		name string
		got string
		want string
	} {
		{
			"Test 576.73",
			problemas.NotasEMoedas(576.73),
			"NOTAS:\n5 nota(s) de R$ 100.00\n1 nota(s) de R$ 50.00\n1 nota(s) de R$ 20.00\n0 nota(s) de R$ 10.00\n1 nota(s) de R$ 5.00\n0 nota(s) de R$ 2.00\nMOEDAS:\n1 moeda(s) de R$ 1.00\n1 moeda(s) de R$ 0.50\n0 moeda(s) de R$ 0.25\n2 moeda(s) de R$ 0.10\n0 moeda(s) de R$ 0.05\n3 moeda(s) de R$ 0.01\n",
		},
		{
			"Test 4.00",
			problemas.NotasEMoedas(4),
			"NOTAS:\n0 nota(s) de R$ 100.00\n0 nota(s) de R$ 50.00\n0 nota(s) de R$ 20.00\n0 nota(s) de R$ 10.00\n0 nota(s) de R$ 5.00\n2 nota(s) de R$ 2.00\nMOEDAS:\n0 moeda(s) de R$ 1.00\n0 moeda(s) de R$ 0.50\n0 moeda(s) de R$ 0.25\n0 moeda(s) de R$ 0.10\n0 moeda(s) de R$ 0.05\n0 moeda(s) de R$ 0.01\n",
		},{
			"Test 91.01",
			problemas.NotasEMoedas(91.01),
			"NOTAS:\n0 nota(s) de R$ 100.00\n1 nota(s) de R$ 50.00\n2 nota(s) de R$ 20.00\n0 nota(s) de R$ 10.00\n0 nota(s) de R$ 5.00\n0 nota(s) de R$ 2.00\nMOEDAS:\n1 moeda(s) de R$ 1.00\n0 moeda(s) de R$ 0.50\n0 moeda(s) de R$ 0.25\n0 moeda(s) de R$ 0.10\n0 moeda(s) de R$ 0.05\n1 moeda(s) de R$ 0.01\n",
		},
	}

	for _, test := range notasemoedastests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.got, test.want)
		})
	}


}


 	