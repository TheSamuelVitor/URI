package main

import (
	"fmt"
	"problemasURI/problemas"
)

func main() {
	var a, b, c string
	fmt.Scanf("%s\n%s\n%s", &a, &b, &c)
	fmt.Println(problemas.Animal(a, b, c))
}
