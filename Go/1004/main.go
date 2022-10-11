package main

import "fmt"

func calculaProduto() {
	var primeiroNumero, segundoNumero int

	fmt.Scanln(&primeiroNumero)
	fmt.Scanln(&segundoNumero)

	fmt.Println("PROD =", primeiroNumero*segundoNumero)
}

func main() {
	calculaProduto()
}