package arreglosslices

import "fmt"

var tablaS []int = []int{2, 5, 4, 6}

var arreglo [10]int = [10]int{1, 52, 36, 96, 87, 54, 12, 32, 59}

func MuestroSlice() {
	fmt.Println(tablaS)

	porcion := arreglo[3:]   //slice desde un vecto, desde la pocision 3 hasta el final
	porcion2 := arreglo[:5]  //slice desde un vecto, desde la pocision 0 hasta la 5
	porcion3 := arreglo[5:7] //slice desde un vecto, desde la pocision 5 hasta la 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)

	fmt.Printf("largo %d, Capacidad %d ", len(elementos), cap(elementos))

	elementos1 := make([]int, 0, 0)

	for i := 0; i < 100; i++ {
		elementos1 = append(elementos1, i)
	}

	fmt.Printf("\nlargo %d, Capacidad %d ", len(elementos1), cap(elementos1))
}
