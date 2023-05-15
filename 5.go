// Escreva um programa que receba uma string e verifique se ela é um número válido em ponto flutuante.
// Informe ao usuário se é ou não.
package main

import (
	"fmt"
	strings2 "strings"
)

func main() {
	var pont string
	fmt.Print("Informe uma string para saber se ela é um número válido em ponto flutuante ou não ")
	fmt.Scan(&pont)
	comparar := strings2.ContainsAny(pont, ".")
	if comparar == true {
		fmt.Print("O número tem ponto flutuante")
	} else {
		fmt.Print("O número não tem ponto flutuante")
	}
}
