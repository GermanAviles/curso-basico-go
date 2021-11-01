package main

import "log"

// "fmt"
// "os"

func main() {

	// archivo := "file-g.txt"

	// f, err := os.Open(archivo)

	// defer f.Close()

	// if err != nil {
	// 	fmt.Println("ALO POLICIA, UN ERROR")
	// 	os.Exit(1)
	// }

	ejemploPanic()
}

func ejemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrió un error, que generó painc \n %v", reco)
		}
	}()

	a := 1

	if a == 1 {
		panic("Se encontró un error")
	}
}
