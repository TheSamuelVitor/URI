package test

import (
	"problemasURI/problemas"
	"testing"
)

func TestSoma(t *testing.T) {
	t.Run("teste soma basico", func(t *testing.T) {
		got := problemas.SomaValores2(5, 5)
		want := 10
		if got != want {
			t.Errorf("\ngot %v, want %v", got, want)
		}
	})
}
