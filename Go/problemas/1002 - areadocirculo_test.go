package problemas

import "testing"

func TestAreadoCirculo(t *testing.T) {
	t.Run("teste area do circulo", func(t *testing.T) {
		got := CalculoAreadoCirculo(2)
		want := "12.5664"
		if got != want {
			t.Errorf("\ngot: %v\nwant: %v", got, want)
		}
	})
}
