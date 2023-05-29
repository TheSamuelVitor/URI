package test

import (
	"problemasURI/assert"
	"problemasURI/problemas"
	"testing"
)

func Test_Media(t *testing.T) {

	tests := []struct {
		name  string
		NotaA float64
		NotaB float64
		NotaC float64
	}{
		{
			"Media maior que 6", 5, 6, 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, problemas.CalculoMedia2(tt.NotaA, tt.NotaB, tt.NotaC), 6.3)
		})
	}

}
