// Escreva um programa que receba uma string e verifique se ela é um palíndromo.
// ]Informe ao usuário se é ou não.
package main

import "fmt"

func main() {
	var s string
	fmt.Print("Infomorme uma string: ")
	fmt.Scan(&s)
	for a := 0; a < len(s)/2; a++ {
		if s[a] != s[len(s)-a-1] {
			fmt.Print("Não é  palidromo")
		} else {
			fmt.Print("é um palidromo")
		}
	}

}
