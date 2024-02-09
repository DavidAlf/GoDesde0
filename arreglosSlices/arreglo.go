package arreglosslices

import "fmt"

var tabla [10]int = [10]int{10, 0, 59}

var matris [4][5]int

func MuestroArreglo() {
	tabla[7] = 33
	tabla[3] = 17

	fmt.Println(tabla)

	tabla2 := [10]string{"Jorge", "David", "Alfonso"}

	for i := 0; i < len(tabla2); i++ {
		fmt.Println(tabla2[i])
	}

	matris[2][2] = 233

	fmt.Println(matris)
}
