package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var texto string

func TablaDeMultiplicar() (int, string) {

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

	texto += fmt.Sprintf("Va a crear la tabla del %d \n", numero)
	for i := 0; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, numero*i)
	}
	texto += "\n"

	return numero, texto
}
