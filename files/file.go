package files

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/DavidAlf/GoDesde0/ejercicios"
)

var numero, texto = ejercicios.TablaDeMultiplicar()
var NombreArchivo string = "./files/txt/TablaDel" + strconv.Itoa(numero) + ".txt"
var ArchivoTotal string = "./files/txt/Tablas.txt"

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
	archivo, err := os.Create(ArchivoTotal)

	if err != nil {
		fmt.Println("Error creando el archivo total ", err.Error())

		return
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {

	if !ExisteArchivo(NombreArchivo) {
		GuardaTabla()
	}

	if !ExisteArchivo(ArchivoTotal) {
		GuardaTablaTotal()
	} else if !Append(ArchivoTotal, texto) {
		fmt.Println("Error al concatenar contenido")
	}

}

func Append(fileNom string, texto string) bool {
	archivo, _ := os.OpenFile(fileNom, os.O_WRONLY|os.O_APPEND, 0644)

	_, err := archivo.WriteString(texto)

	if err != nil {
		fmt.Println("Error realizando el WriteString ", err.Error())

		return false
	}
	archivo.Close()

	return true
}

func ExisteArchivo(ruta string) bool {
	if _, err := os.Stat(ruta); os.IsNotExist(err) {
		return false
	}
	return true
}

func LeoArchivo() {
	archivo, err := os.Open(NombreArchivo)

	if err != nil {
		fmt.Println("Error leyendo el archivo " + err.Error())

		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()

		fmt.Println("> " + registro)
	}

	archivo.Close()

}
