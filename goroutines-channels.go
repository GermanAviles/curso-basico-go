package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	canal1 := make(chan time.Duration)
	go nombreLento("Germán David", canal1)
	fmt.Println("Punto 12")
	valor := <-canal1
	fmt.Printf("Tiempo espera: %v", valor)
}

func nombreLento(nombre string, canal chan time.Duration) {
	tiempoInicio := time.Now()
	letras := strings.Split(nombre, "")

	for _, letra := range letras {
		// time.Sleep(1 * time.Second)
		fmt.Printf(letra)
	}

	tiempoFin := time.Now()
	duracion := tiempoFin.Sub(tiempoInicio)
	fmt.Printf("\n")
	// fmt.Println(duracion)
	canal <- duracion
}
