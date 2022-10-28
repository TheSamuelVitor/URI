package problemas

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("teste Hello world", func(t *testing.T) {
		got := MostraMensagem()
		want := "Hello World!"
		if got != want {
			t.Errorf("\ngot: %v\nwant: %v", got, want)
		}
	})
}