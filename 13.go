// Solicite ao usuário uma string e
// informe se ela é uma sequência numérica crescente (exemplo: "123" ou "4567").
package main

import "fmt"

func main() {
	var x string
	fmt.Print("Informe uma string")
	fmt.Scan(&x)
	c := true
	for i := 0; i < len(x)-1; i++ {
		if x[i] >= x[i+1] {
			c = false
		}
	}
	if c {
		fmt.Print("Esta em ordem crescente")
	} else {
		fmt.Print("Não esta em ordem crescente")
	}
}
