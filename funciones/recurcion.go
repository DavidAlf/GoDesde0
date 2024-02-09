package funciones

import (
	"fmt"
	"strconv"
)

func Exponencia(valor int) {
	if valor > 100000000 {
		return
	}

	fmt.Println("Valor " + strconv.Itoa(valor))
	Exponencia(valor * 2)
}
