package files

import (
	"fmt"
	"os"
	"strconv"

	"github.com/DavidAlf/GoDesde0/ejercicios"
)

var numero, texto = ejercicios.TablaDeMultiplicar()
var NombreArchivo string = "./files/txt/TablaDel" + strconv.Itoa(numero) + ".txt"
var total string = "./files/txt/Tablas.txt"

func GuardaTabla() {
	archivo, err := os.Create(NombreArchivo)

	if err != nil {
		fmt.Println("Error creando el archivo ", err.Error())

		return
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func GuardaTablaTotal() {
	archivo, err := os.Create(total)

	if err != nil {
		fmt.Println("Error creando el archivo total ", err.Error())

		return
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {

	if Append(total, texto) == false {
		fmt.Println("Error al concatenar contenido")
	}

}

func Append(fileNom string, texto string) bool {
	GuardaTabla()

	arch, err := os.OpenFile(fileNom, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		GuardaTablaTotal()

		arch, err = os.OpenFile(fileNom, os.O_WRONLY|os.O_APPEND, 0644)
	}

	_, err = arch.WriteString(texto)

	if err != nil {
		fmt.Println("Error realizando el WriteString ", err.Error())

		return false
	}

	arch.Close()

	return true
}
