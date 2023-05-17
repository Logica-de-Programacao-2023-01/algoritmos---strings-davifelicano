package main

import (
	"fmt"
)

func main() {
	var x string
	fmt.Print("Informe uma string: ")
	fmt.Scan(&x)

	isDecrescente := true
	for i := 0; i < len(x)-1; i++ {
		if x[i] < x[i+1] {
			isDecrescente = false
			break
		}
	}

	if isDecrescente {
		fmt.Println("A string está em ordem decrescente")
	} else {
		fmt.Println("A string não está em ordem decrescente")
	}
}
