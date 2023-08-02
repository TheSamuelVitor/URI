package problemas

// Lê o valor de 3 notas e calcula a média entre elas
// Sendo que a primeira nota vale 2, a segunda 3 e a terceira 5
func CalculoMedia2(notaA, notaB, notaC float64) (media float64) {

	media = ((notaA * 2) + (notaB * 3) + (notaC * 5)) / 10

	return media
}
