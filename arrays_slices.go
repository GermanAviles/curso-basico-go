package main

import "fmt"

// var tabla [10]int8
// var matriz [5][7]int

func main() {
	fmt.Println("=== Arrays ===")
	// matriz[3][5] = 1
	tabla := [10]int{1, 2, 3, 4, 5, 6, 7, 8}
	// tabla[0] = 1
	// tabla[1] = 10
	fmt.Println(tabla)
	// fmt.Println(matriz)
	fmt.Println("=== Slices ===")
	slice := []int{1, 2, 3, 4}
	fmt.Println(slice)
	slice2()
	slice3()
	slice4()
}

func slice2() {
	elementos := [6]int{1, 2, 3, 4, 5, 6}
	// slice := elementos[:3]
	// slice := elementos[3:]
	slice := elementos[2:5]
	fmt.Println(slice)
}

func slice3() {
	slice := make([]int, 5, 20)
	fmt.Printf("Largo: %d, Capacidad: %d \n", len(slice), cap(slice))
}

func slice4() {
	slice := make([]int, 0)
	for i := 0; i < 100; i++ {
		slice = append(slice, i)
	}
	fmt.Printf("Largo: %d, Capacidad: %d \n", len(slice), cap(slice))
}
