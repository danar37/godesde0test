package main

import (
	"fmt"
	"runtime"
)

func main() {
	/*estado, texto := variables.ConviertoaTexto(2500)
	fmt.Println(estado)
	fmt.Println(texto) */

	if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("Esto no es Windows, es ", os)
	} else {
		fmt.Println("Esto es Windows")
	}

}
