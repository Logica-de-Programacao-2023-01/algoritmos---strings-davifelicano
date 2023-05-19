// Solicite ao usuário uma string e imprima a mesma string invertida.
package main

import "fmt"

func main() {
	var x string
	fmt.Print("Informe uma string: ")
	fmt.Scan(&x)
	j := []rune(x)
	t := len(j)
	for d, k := 0, t-1; d < k; d, k = d+1, k-1 {
		j[d], j[k] = j[k], j[d]
	}
	trocar := string(j)
	fmt.Print("trocado:  ", trocar)
}
