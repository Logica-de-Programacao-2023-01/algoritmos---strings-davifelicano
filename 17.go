// Solicite ao usuário uma string e imprima somente as suas letras únicas (que aparecem apenas uma vez).
package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Print("Escreva uma string:")
	fmt.Scan(&x)
	for _, a := range x {
		contador := strings.Count(x, string(a))
		if contador == 1 {
			fmt.Print(string(a))
		}
	}
}
