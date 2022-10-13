package problemas

import "fmt"

func OMaior() {
	var valor1, valor2, valor3 int

	fmt.Scanf("%d %d %d", &valor1, &valor2, &valor3)

	if valor1 >= valor2 && valor1 >= valor3 {
		fmt.Println(valor1, "eh o maior")
	} else if valor2 >= valor1 && valor2 >= valor3 {
		fmt.Println(valor2, "eh o maior")
	} else if valor3 >= valor2 && valor3 >= valor1 {
		fmt.Println(valor3, "eh o maior")
	}
}
