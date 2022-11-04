package problemas

import "testing"

func TestIntervalo(t *testing.T) {

	intervaloTests := []struct {
		name string
		got string
		want string
	} {
		{ name: "test 0 a 25", got: Intervalo(25.00), want: "Intervalo [0,25]"},
		{ name: "test 25 a 50", got: Intervalo(26.00), want: "Intervalo (25,50]"},
		{ name: "test 50 a 75", got: Intervalo(55.00), want: "Intervalo (50,75]"},
		{ name: "test 75 a 100", got: Intervalo(25.00), want: "Intervalo [0,25]"},
		{ name: "test fora de intervalo", got: Intervalo(-25.00), want: "Fora de intervalo"},
	}

	for _, test := range intervaloTests {
		t.Run(test.name, func(t *testing.T){
			got := test.got
			want := test.want
			if got != want {
				t.Errorf("\ngot %v\nwant %v", got, want)
			}
		})
	}

}
