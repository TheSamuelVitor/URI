package test

import (
	"problemasURI/assert"
	"problemasURI/problemas"
	"testing"
)

func TestDeSelecao(t *testing.T) {

	selecaoTests := []struct {
		name string
		got  string
		want string
	}{
		{
			"Teste Valores Nao Aceitos",
			problemas.TesteDeSelecao(5, 6, 7, 8),
			"Valores nao aceitos",
		}, {
			"Teste Valores Aceitos",
			problemas.TesteDeSelecao(2, 3, 2, 6),
			"Valores aceitos",
		},
	}

	for _, test := range selecaoTests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.got, test.want)
		})
	}

}
