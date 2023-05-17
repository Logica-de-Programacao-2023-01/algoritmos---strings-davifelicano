// Escreva um programa que receba duas strings e
// verifique se elas são anagramas. Informe ao usuário se são ou não.
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var x1, x2 string
	fmt.Print("Informe uma string para que seja verificado se ela é um anagrama com outra string:  ")
	fmt.Scan(&x1)
	fmt.Print("Informe outra string para seja comparada com a primeira : ")
	fmt.Scan(&x2)
	if len(x1) != len(x2) {
		fmt.Print("Não é uma anagrama ,pois o tamanho é diferente")
		return
	}
	s1 := strings.ReplaceAll(strings.ToLower(x1), " ", "")
	s2 := strings.ReplaceAll(strings.ToLower(x2), " ", "")
	st1 := strings.Split(s1, "")
	st2 := strings.Split(s2, "")
	sort.Strings(st1)
	sort.Strings(st2)
	if strings.Join(st1, "") == strings.Join(st2, "") {
		fmt.Print("Estas strings são anagramas ")
	}
}
