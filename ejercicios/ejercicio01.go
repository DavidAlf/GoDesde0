package ejercicios

import "strconv"

func ValidaNumero(dato string) (numero int, respuesta string) {

	numero, _ = strconv.Atoi(dato)

	if numero > 100 {
		respuesta = "El numero es mayor a 100"
	} else {
		respuesta = "El numero es menor a 100"
	}

	return numero, respuesta
}
