package test

import (
	"problemasURI/assert"
	"problemasURI/problemas"
	"testing"
)

func TestProduto(t *testing.T) {

	produtoTests := []struct{
		name string
		got string
		want string
	}{
		{name: "Teste Funcional 1", got: problemas.CalculaProduto(5,7), want: "PROD = 35"},
	}

	for _, test := range produtoTests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.got, test.want)
		})
	}

}
