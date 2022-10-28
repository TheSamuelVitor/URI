package problemas

import "testing"

func TestIntervalo(t *testing.T) {

	// t.Run("test 0 a 25", func(t *testing.T) {
	// 	got := Intervalo(25.00)
	// 	want := "Intervalo [0,25]"
	// 	if got != want {
	// 		t.Errorf("\ngot %v\nwant %v", got, want)
	// 	}
	// })

	// t.Run("test 25 a 50", func(t *testing.T) {
	// 	got := Intervalo(25.01)
	// 	want := "Intervalo (25,50]"
	// 	if got != want {
	// 		t.Errorf("\ngot %v\nwant %v", got, want)
	// 	}
	// })

	// t.Run("test 50 a 75", func(t *testing.T) {
	// 	got := Intervalo(55.01)
	// 	want := "Intervalo (50,75]"
	// 	if got != want {
	// 		t.Errorf("\ngot %v\nwant %v", got, want)
	// 	}
	// })

	// t.Run("test 75 a 100", func(t *testing.T) {
	// 	got := Intervalo(100.00)
	// 	want := "Intervalo (75,100]"
	// 	if got != want {
	// 		t.Errorf("\ngot %v\nwant %v", got, want)
	// 	}
	// })

	// t.Run("test fora de intervalo", func(t *testing.T) {
	// 	got := Intervalo(101.00)
	// 	want := "Fora de intervalo"
	// 	if got != want {
	// 		t.Errorf("\ngot %v\nwant %v", got, want)
	// 	}
	// })

	intervaloTests := []struct {
		name string
		got string
		want string
	} {
		{ name: "test 0 a 25", got: Intervalo(25.00), want: "Intervalo [0,25]"},
		{ name: "test 25 a 50", got: Intervalo(26.00), want: "Intervalo [0,25]"},
		{ name: "test 50 a 75", got: Intervalo(55.00), want: "Intervalo [0,25]"},
		{ name: "test 75 a 100", got: Intervalo(25.00), want: "Intervalo [0,25]"},
		{ name: "test fora de intervalo", got: Intervalo(25.00), want: "Fora de intervalo"},
	}

}
