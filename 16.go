// Solicite ao usuário duas strings e informe se a segunda é uma substring da primeira.
package main

import (
	"fmt"
	"strings"
)

func main() {
	var x, y string
	fmt.Print("Informe uma string: ")
	fmt.Scan(&x)
	fmt.Print("Informe uma string: ")
	fmt.Scan(&y)
	contais := strings.Contains(x, y)
	if contais {
		fmt.Print("È uma substring")
	} else {
		fmt.Print("Não é uma substring")
	}

}
