// Solicite ao usuário duas strings e informe se elas são iguais ou diferentes.
package main

import (
	"fmt"
	"strings"
)

func main() {
	var string1 string
	var strng2 string
	fmt.Println("Escreva uma string ")
	fmt.Scan(&string1)
	fmt.Println("Escreva outra string")
	fmt.Scan(&strng2)
	fmt.Println("As strings serão comparadas")
	comparar := strings.Compare(string1, strng2)
	if comparar == 0 {
		fmt.Print("As strings são iguais")
	} else {
		fmt.Print("As strings são diferentes")
	}

}
