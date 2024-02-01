package ejercicios

import "strconv"

func ValidaNumero(dato string) (int, string) {

	numero, error := strconv.Atoi(dato)

	if error != nil {
		return numero, "Hubo un error" + error.Error()
	}

	if numero > 100 {
		return numero, "El numero es mayor a 100"
	} else {
		return numero, "El numero es menor a 100"
	}
}
