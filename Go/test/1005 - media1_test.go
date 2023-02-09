package test

import (
	"problemasURI/assert"
	"problemasURI/problemas"
	"testing"
)

func TestMedia1(t *testing.T) {

	media1Tests := []struct {
		name string
		got  string
		want string
	}{
		{
			"Teste Media 1",
			problemas.CalculoMedia(7, 8),
			"MEDIA = 7.68182",
		},
	}

	for _, test := range media1Tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.got, test.want)
		})
	}

}
