package test

import (
	"problemasURI/assert"
	"problemasURI/problemas"
	"testing"
)

func TestTiposDeTriangulo(t *testing.T) {

	tiposDeTriangulotests := []struct {
		name   string
		got    string
		expect string
	}{
		{
			"Teste 1",
			problemas.TiposDeTriangulo(7, 5, 7),
			"TRIANGULO ACUTANGULO\nTRIANGULO ISOSCELES",
		},
		{
			"Teste 2",
			problemas.TiposDeTriangulo(6, 6, 10),
			"TRIANGULO OBTUSANGULO\nTRIANGULO ISOSCELES",
		},
		{
			"Teste 3",
			problemas.TiposDeTriangulo(6, 6, 6),
			"TRIANGULO ACUTANGULO\nTRIANGULO EQUILATERO",
		},
		{
			"Teste 4",
			problemas.TiposDeTriangulo(7, 5, 2),
			"NAO FORMA TRIANGULO",
		},
	}

	for _, test := range tiposDeTriangulotests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, test.got)
		})
	}

}
