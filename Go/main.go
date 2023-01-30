package main

import (
	"fmt"
	"problemasURI/problemas"
)

func main() {
	var coordenadaX, coordenadaY float64
	fmt.Scanf("%v %v\n", &coordenadaX, &coordenadaY)
	resultado := problemas.CoordenadasDeUmPonto(coordenadaX, coordenadaY)
	fmt.Printf("%s\n", resultado)
}
