package variables

import (
	"fmt"
	"strconv"
	"time"
)

var nombre string
var estado bool
var sueldo float32
var fecha time.Time

func RestoVariables() {
	nombre = "David"
	estado = true
	sueldo = 6000.5
	fecha = time.Now()

	fmt.Println("Nombre = ", nombre)
	fmt.Println("Estado = ", estado)
	fmt.Println("Sueldo = ", sueldo)
	fmt.Println("Fecha = ", fecha)

}

func ConviertoATexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)

	return true, texto
}
