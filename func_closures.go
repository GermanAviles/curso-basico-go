package main

import "fmt"

var calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {

	fmt.Printf("5 + 7: %d \n", calculo(5, 7))

	calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Printf("6 - 4: %d \n", calculo(6, 4))

	operaciones()
	/*=== Closures === */
	tablaDel := 2
	miTabla := tabla(tablaDel)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d \n", tablaDel, i, miTabla())
	}
}

func operaciones() {
	resultado := func() int {
		var a = 23
		var b = 13
		return a + b
	}

	fmt.Println(resultado())
}

func tabla(value int) func() int {
	numero := value
	secuencia := 0

	return func() int {
		secuencia += 1
		return numero * secuencia
	}

}
