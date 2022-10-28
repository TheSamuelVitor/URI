package problemas

import "testing"

func TestDeSelecao(t *testing.T) {
	t.Run("valores nao aceitos", func(t *testing.T) {
		got := TesteDeSelecao(5,6,7,8)
		want := "Valores nao aceitos"
		if got != want {
			t.Errorf("\ngot %v\nwant %v", got, want)
		}
	})

	t.Run("valores aceitos", func(t *testing.T) {
		got := TesteDeSelecao(2,3,2,6)
		want := "Valores aceitos"
		if got != want {
			t.Errorf("\ngot %v\nwant %v", got, want)
		}
	})
}