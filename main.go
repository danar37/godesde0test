package main

import (
	"fmt"

	"github.com/danar37/godesde0test/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1578)
	fmt.Println(estado)
	fmt.Println(texto)
	fmt.Println()
}
