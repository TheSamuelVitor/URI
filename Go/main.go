package main

import (
	"fmt"
	"problemasURI/problemas"
)

func main() {
	var n1, n2, n3 int
	fmt.Scanf("%v %v %v\n", &n1, &n2, &n3)
	resultado := problemas.SortSimples(n1, n2, n3)
	fmt.Printf("%s\n", resultado)
}
