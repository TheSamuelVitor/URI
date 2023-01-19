package problemas

import (
	"problemasURI/assert"
	"testing"
)

func TestAreadoCirculo(t *testing.T) {
	t.Run("teste area do circulo", func(t *testing.T) {
		got := CalculoAreadoCirculo(2)
		want := "12.5664"
		assert.Equal(t, got, want)
	})
}
