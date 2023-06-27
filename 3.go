// Escreva um programa que receba uma string e
// substitua todas as ocorrências desse caractere na string por outro caractere especificado pelo usuário.
// Informe ao usuário o resultado.
package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra string
	var tirar string
	var colocar string
	fmt.Print("Escreva uma string ")
	fmt.Scan(&palavra)
	fmt.Print("Escreva uma letra para substituir ")
	fmt.Scan(&tirar)
	fmt.Print("Escreva uma nova letra ")
	fmt.Scan(&colocar)
	novastring := strings.ReplaceAll(palavra, tirar, colocar)
	fmt.Print(novastring)
}
