package mapas

import "fmt"

func MostratMapas() {
	paises := make(map[string]string)
	fmt.Println(paises)

	paises["Mexico"] = "D.F"
	paises["Colombia"] = "Bogota"
	paises["Argentina"] = "Buenos Aires"

	/*	fmt.Println(paises)
		fmt.Println("")
		fmt.Println(paises["Colombia"])
	*/
	campeonato := map[string]int{
		"Barcelona":   39,
		"Real Madrid": 38,
		"Chivas":      12,
		"Boca Junior": 20,
	}

	fmt.Println(campeonato)
	/*fmt.Println("")
	fmt.Println(campeonato["Chivas"])
	*/
	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	}

	delete(campeonato, "Chivas")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Barcelona"]

	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)
}
