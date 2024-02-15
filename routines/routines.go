package routines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLento(nombre string) {
	letras := strings.Split(nombre, "")

	for _, letra := range letras {
		time.Sleep(1 * time.Second)

		fmt.Println(letra)
	}

}
