package problemas

import "testing"

func TestBasico(t *testing.T) {
	t.Run("testando soma de dois valores", func(t *testing.T) {
		got := SomaValores(5, 5)
		want := 10
		if got != want {
			t.Errorf("\ngot: %v\nwant: %v", got, want)
		}
	})
}
