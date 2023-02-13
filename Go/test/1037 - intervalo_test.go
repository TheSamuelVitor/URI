package test

import (
	"problemasURI/assert"
	"problemasURI/problemas"
	"testing"
)

func TestIntervalo(t *testing.T) {

	intervaloTests := []struct {
		name string
		got  string
		want string
	}{
		{name: "test 0 a 25", got: problemas.Intervalo(25.00), want: "Intervalo [0,25]"},
		{name: "test 25 a 50", got: problemas.Intervalo(26.00), want: "Intervalo (25,50]"},
		{name: "test 50 a 75", got: problemas.Intervalo(55.00), want: "Intervalo (50,75]"},
		{name: "test 75 a 100", got: problemas.Intervalo(25.00), want: "Intervalo [0,25]"},
		{name: "test fora de intervalo", got: problemas.Intervalo(-25.00), want: "Fora de intervalo"},
	}

	for _, test := range intervaloTests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.got, test.want)
		})
	}

}
