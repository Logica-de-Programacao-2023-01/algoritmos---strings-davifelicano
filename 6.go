// Escreva um programa que receba uma string e conte quantas palavras ela contém.
// Informe ao usuário o resultado.
package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Print("Informe um número: ")
	fmt.Scanln(&x)
	a := strings.Fields(x)

	for _, c := range a {
		fmt.Print(len(c) + 1)
	}
}
