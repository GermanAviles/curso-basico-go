package main

import "fmt"

func main() {
	paises := make(map[string]string, 5)

	paises["COL"] = "Colombia"
	paises["ARG"] = "Argentina"
	fmt.Println(paises)
	fmt.Println(paises["COL"])

	campeonatos := map[string]int{"COL": 1, "ARG": 2}
	campeonatos["PER"] = 2
	campeonatos["COL"] = 20
	campeonatos["MEX"] = 10
	fmt.Println(campeonatos)
	delete(campeonatos, "PER")
	fmt.Println(campeonatos)

	for equipo, puntaje := range campeonatos {
		fmt.Printf("Equipo: %s - Puntos: %d \n", equipo, puntaje)
	}
	valor, existe := campeonatos["LOL"]
	fmt.Printf("Son: %d puntos - el equpo existe: %t \n", valor, existe)
	// if existe {
	// }
}
