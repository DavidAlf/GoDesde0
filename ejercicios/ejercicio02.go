package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var texto string
var err error

func TablaDeMultiplicar() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese el numero")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())

			if err != nil {
				fmt.Println("Error con el numero ingresado " + err.Error())
				continue
			}
			break
		}
	}

	fmt.Println("Va a crear la tabla de multiplicar", numero)
	for i := 0; i <= 10; i++ {
		texto = strconv.Itoa(numero) + " * " + strconv.Itoa(i) + " = "
		fmt.Println(texto, numero*i)
	}
}
