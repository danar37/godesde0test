package main

import (
	"fmt"

	"github.com/danar37/godesde0test/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(2500)
	fmt.Println(estado)
	fmt.Println(texto)
}
