// Solicite ao usuário uma string e informe se ela é uma sequência numérica decrescente (exemplo: "987" ou "54321").
package main

import "fmt"

func main() {
	var x string
	fmt.Print("Informe uma string para verificar se ela é decrescente")
	fmt.Scan(&x)
	d := true
	for j := 0; j < len(x)-1; j++ {
		if x[j] <= x[j+1] {
			d = false
		}
	}
	if d {
		fmt.Print("Esta em ordem descrecente")
	} else {
		fmt.Print("não esta em ordem decrescente")
	}
}
