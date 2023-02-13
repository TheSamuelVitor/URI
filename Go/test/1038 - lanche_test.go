package test

import (
	"problemasURI/assert"
	"problemasURI/problemas"
	"testing"
)

func TestLanche(t *testing.T) {

	lancheTests := []struct {
		name string
		got  float64
		want float64
	}{
		{name: "Teste 2,3", got: problemas.Lanche(2, 3), want: 13.5},
		{name: "Teste 3,2", got: problemas.Lanche(3, 2), want: 10},
		{name: "Teste 4,3", got: problemas.Lanche(4, 3), want: 6},
		{name: "Teste 1,4", got: problemas.Lanche(1, 4), want: 16},
		{name: "Teste 5,1", got: problemas.Lanche(5, 1), want: 1.5},
	}

	for _, test := range lancheTests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.got, test.want)
		})
	}
}
