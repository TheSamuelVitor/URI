package test

import (
	"problemasURI/assert"
	"problemasURI/problemas"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("teste Hello world", func(t *testing.T) {
		got := problemas.MostraMensagem()
		want := "Hello World!"
		assert.Equal(t, got, want)
	})
}
