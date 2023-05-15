// Solicite ao usuário uma string e informe se ela contém pelo menos um número.
package main

import (
	"fmt"
	"unicode"
)

func main() {
	var x string
	fmt.Print("Informe uma string: ")
	fmt.Scan(&x)
	temn := false
	for _, c := range x {
		if unicode.IsNumber(c) {
			temn = true
			break
		}
	}
	if temn {
		fmt.Print("tem número")

	} else {
		fmt.Print("Não tem número")
	}

}
