package problemas

import "fmt"

func SortSimples(n1, n2, n3 int) (resultado string) {
	var numero1, numero2, numero3 int

	if n1 > n2 && n1 > n3 {
		numero1 = n1
		if n2 > n3 {
			numero2 = n2
			numero3 = n3
		} else {
			numero2 = n3
			numero3 = n2
		}
	} else if n2 > n1 && n2 > n3 {
		numero1 = n2
		if n1 > n3 {
			numero2 = n1
			numero3 = n3
		} else {
			numero2 = n3
			numero3 = n1
		}
	} else if n3 > n2 && n3 > n1 {
		numero1 = n3
		if n1 > n2 {
			numero2 = n1
			numero3 = n2
		} else {
			numero2 = n2
			numero3 = n1
		}
	}

	resultado = fmt.Sprintf("%d\n%d\n%d\n\n%d\n%d\n%d", numero3, numero2, numero1, n1, n2, n3)

	return

}