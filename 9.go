// Solicite ao usuário uma string e substitua todas as ocorrências de uma letra por outra informada pelo usuário.
package main

import (
	"fmt"
	"strings"
)

func main() {
	var x, x1, x3 string
	fmt.Println("Informe uma string:")
	fmt.Scanln(&x)
	fmt.Println("Informe um caracter que deseja ser substituido:")
	fmt.Scanln(&x1)
	fmt.Println("Informe o novo caracter:")
	fmt.Scanln(&x3)
	NovaString := strings.ReplaceAll(x, x1, x3)
	fmt.Println("Sua nova string:", NovaString)
}
