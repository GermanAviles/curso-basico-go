package main

import "fmt"

func primeraFuncion(valor int) int {
	fmt.Printf("VALOR: %d \n", valor)
	return valor + 10
}

func segundaFuncion(valor int) (string, bool) {
	fmt.Printf("VALOR: %d \n", valor)

	if valor%2 == 0 {
		return "Es par carnal!!!!", true
	}

	return "Es impar carnal :c", false
}

func calculo(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total = total + numero
	}
	return total
}

func main() {
	// nuevoValor := primeraFuncion(190)
	// valuePar, esPar := segundaFuncion(1)
	sumatoria := calculo(10, 20)
	// fmt.Printf("Nuevo valor %d \n", nuevoValor)
	// fmt.Printf("%v \n", valuePar)
	// fmt.Printf("es PAR: %v", esPar)
	fmt.Printf("sumatoria: %v", sumatoria)

}
