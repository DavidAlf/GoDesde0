package middleware

import (
	"fmt"
)

func sumar(a, b int) int {

	return a + b
}

func restar(a, b int) int {

	return a - b
}

func multiplicar(a, b int) int {

	return a * b
}

func MiMiddleware() {
	fmt.Println("Inicio")

	result := operacionesMiddleware(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMiddleware(restar)(2, 3)
	fmt.Println(result)
	result = operacionesMiddleware(multiplicar)(2, 3)
	fmt.Println(result)

	fmt.Println("Fin")
}

func operacionesMiddleware(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operacion")

		return f(a, b)
	}
}
