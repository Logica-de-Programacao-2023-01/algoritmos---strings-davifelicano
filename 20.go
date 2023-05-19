// Solicite ao usuário uma string e informe se ela é está em camelCase
// e quantas palavras possuí. CamelCase é quando a primeira letra de cada palavra é maiúscula,
// e as demais são minúsculas. Exemplo: "estaStringEstaEmCamelCase".
package main

import (
	"fmt"
	"strings"
)

func main() {
	var j string
	p := 0
	fmt.Print("Informe uma string")
	fmt.Scan(&j)
	if strings.ToLower(string(j[0])) != string(j[0]) {
		fmt.Println("Esta esta em CamelCase")
	}
	for _, x := range j {
		if strings.ToUpper(string(x)) == string(x) {
			p++
		}
	}

	fmt.Println("Quantidade de palavras em CamelCase: ", p)
}
