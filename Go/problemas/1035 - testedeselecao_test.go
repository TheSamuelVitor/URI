package problemas

import "testing"

func TestDeSelecao(t *testing.T) {
	t.Run("valores nao aceitos", func(t *testing.T) {
		got := "o"
		want := "Valores nao aceitos"
		if got != want {
			t.Errorf("\ngot %v\nwant %v", got, want)
		}
	})

	t.Run("valores aceitos", func(t *testing.T) {
		got := "oi"
		want := "Valores aceitos"
		if got != want {
			t.Errorf("\ngot %v\nwant %v", got, want)
		}
	})
}