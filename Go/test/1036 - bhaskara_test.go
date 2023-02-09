package test

import (
	"problemasURI/assert"
	"problemasURI/problemas"
	"testing"
)

func TestBhaskara(t *testing.T) {

	bhaskaraTests := []struct {
		name, got, want string
	}{
		{
			"Teste Impossivel Calcular",
			problemas.FormuladeBhaskara(0, 20, 5),
			"Impossivel calcular\n",
		},
		{
			"Teste Bhaskara 1",
			problemas.FormuladeBhaskara(10, 20.1, 5.1),
			"R1 = -0.29788\nR2 = -1.71212\n",
		},
		{
			"Teste Bhaskara 2",
			problemas.FormuladeBhaskara(10.3, 203, 5),
			"R1 = -0.02466\nR2 = -19.68408\n",
		},
	}

	for _, test := range bhaskaraTests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.got, test.want)
		})
	}
}
