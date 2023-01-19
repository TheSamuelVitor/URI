package problemas

import (
	"problemasURI/assert"
	"testing"
)

func TestBasico(t *testing.T) {
	t.Run("testando soma de dois valores", func(t *testing.T) {
		got := SomaValores(5, 5)
		want := 10
		assert.Equal(t, got, want)
	})
}
