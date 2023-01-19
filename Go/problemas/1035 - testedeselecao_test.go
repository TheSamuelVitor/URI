package problemas

import (
	"problemasURI/assert"
	"testing"
)

func TestDeSelecao(t *testing.T) {
	t.Run("valores nao aceitos", func(t *testing.T) {
		got := "o"
		want := "Valores nao aceitos"
		assert.Equal(t, got, want)
	})

	t.Run("valores aceitos", func(t *testing.T) {
		got := "oi"
		want := "Valores aceitos"
		assert.Equal(t, got, want)
	})
}