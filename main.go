package main

import (
	"fmt"

	"github.com/DavidAlf/GoDesde0/ejercicios"
)

func main() {
	//	fmt.Println("Variables Enteras")
	//	variables.MuestroEnteros()
	//
	//	fmt.Println()
	//
	//	fmt.Println("Variables Resto")
	//	variables.RestoVariables()
	//
	//	fmt.Println()
	//
	//	fmt.Println("Funcion convierte numro a texto")
	//	estado, texto := variables.ConviertoATexto(1991)
	//	fmt.Println(estado)
	//	fmt.Println(texto)
	//
	//	fmt.Println()
	//
	//	fmt.Println("Sistema operativo")
	//	otros.DatosSistemaOperativos()
	//
	//	fmt.Println()

	fmt.Println("Ejercicio 01")
	numero, respuesta := ejercicios.ValidaNumero("xx")
	fmt.Println(numero)
	fmt.Println(respuesta)
}
