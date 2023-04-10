// Escreva um programa que receba uma string
// e converta todas as letras minúsculas para maiúsculas. Informe ao usuário o resultado.
package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Print("Informe uma string : ")
	fmt.Scan(&x)
	fmt.Print(strings.ToUpper(x))
}
