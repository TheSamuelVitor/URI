package problemas

import (
	"problemasURI/assert"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("teste Hello world", func(t *testing.T) {
		got := MostraMensagem()
		want := "Hello World!"
		assert.Equal(t, got, want)
	})
}
