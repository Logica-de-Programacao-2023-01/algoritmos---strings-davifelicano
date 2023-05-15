// Escreva um programa que receba uma string e remova todos os espaços em branco.
// Informe ao usuário o resultado.
package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Print("Digite a string: ")
	fmt.Scanln(&a)
	b := strings.ReplaceAll(a, " ", "")
	fmt.Println(b)
}
