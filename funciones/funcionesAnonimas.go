package funciones

import "fmt"

func Calculos() {

	var num3 int = 50

	var num4 int = 250

	calculo := func(num1 int, num2 int) int {
		return num1 + num2
	}

	fmt.Println(calculo(10, 25) + num3 + num4)

	calculo = func(num1 int, num2 int) int {
		return num1*num2*num3 - num4
	}

	fmt.Println(calculo(10, 25) + num3 - num4)
}
