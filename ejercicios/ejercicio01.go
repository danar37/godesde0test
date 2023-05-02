package ejercicios

import "strconv"

func Ejemplo(texto) (int, string) {

	num := strconv.Atoi(texto)
	if num > 100 {
		return num, "Es mayor a 100 "
	} else {
		return num, "Es menor a 100"
	}

}
