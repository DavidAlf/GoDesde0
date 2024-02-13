package errores

import (
	"fmt"
	"log"
)

// ultima instruccion cuando termine un proceso
func VemosDefer() {

	fmt.Println("Primer Mensaje")

	defer fmt.Println("Este es el mensaje final")

	fmt.Println("Segundo Mensaje")
}

func EjemploPanic() {

	defer func() {
		reco := recover()

		if reco != nil {
			log.Fatalf("Ocurrio un error que genero PANIC \n %v", reco)
		}
	}()

	a := 1

	if a == 1 {
		panic("Se encontro el valor 1")
	}
}
