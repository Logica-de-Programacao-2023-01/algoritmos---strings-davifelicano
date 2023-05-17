// Solicite ao usuário uma string e substitua todas as vogais por '*' (asterisco).
package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, r string
	fmt.Print("Escreva uma string")
	fmt.Scan(&s)
	vogal := "AEIOUaeiou"
	for _, v := range vogal {

		r = strings.ReplaceAll(s, string(v), "*")
		s = r
	}
	fmt.Print(r)
}
