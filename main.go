package main

import (
	"fmt"
	"runtime"

	"github.com/DavidAlf/GoDesde0/variables"
)

func main() {
	fmt.Println("Variables Enteras")
	variables.MuestroEnteros()

	fmt.Println()

	fmt.Println("Variables Resto")
	variables.RestoVariables()

	fmt.Println()

	fmt.Println("Funcion convierte numro a texto")
	estado, texto := variables.ConviertoATexto(1991)
	fmt.Println(estado)
	fmt.Println(texto)

	fmt.Println()

	fmt.Println("Sistema operativo")

	if os := runtime.GOOS; os == "Linux." {
		fmt.Println("Es Linux")
	} else {
		fmt.Println("Es Windows")
	}
}
