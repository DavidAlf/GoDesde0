package ejer_interfaces

import (
	"fmt"

	"github.com/DavidAlf/GoDesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	fmt.Printf("Soy un/a %s, estoy vivo %t, y estoy respirando\n", hu.Sexo(), hu.EstaVivo())
}
