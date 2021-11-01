package main

import "fmt"

func GoToRutina() {
	var i = 0
RUTINA:
	fmt.Println("Antes del FOR")
	for i < 10 {
		if i == 4 {
			i = i + 2
			fmt.Println("ACTIVAR LA ROUTINEEE")
			goto RUTINA
		}
		fmt.Printf("Valor de i: %d \n", i)
		i++
	}
}

func forContinue() {
	for i := 1; i < 10; i++ {
		if i == 4 {
			fmt.Println("Continue activado")
			continue
		}
		fmt.Printf("Valor de i: %d \n", i)
	}
}

func forBreak() {
	for i := 1; i < 10; i++ {
		if i == 4 {
			fmt.Println("break activado")
			break
		}
		fmt.Printf("Valor de i: %d \n", i)
	}
}

func forInfinite() {
	i := 1
	continuar := true

	for {
		if i == 5 {
			fmt.Println("break activado")
			break
		}
		fmt.Printf("Valor de i: %d \n", i)
		i++
	}

	i = 1
	fmt.Println("==== For con condicion (While) ====")
	for continuar {
		if i == 5 {
			continuar = false
		}
		fmt.Printf("Valor de i: %v \n", continuar)
		i++
	}
}

func main() {
	// GoToRutina()
	// forContinue()
	// forBreak()
	forInfinite()
}
