package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	leoArchivo()
	leoArchivo2()
	graboArchivo()
	graboArchivo2()

}

func leoArchivo() {
	archivo, err := ioutil.ReadFile("./file-go.txt")

	if err != nil {
		// error := fmt.Errorf(err.Error())
		fmt.Println(err)
	}

	fmt.Println(string(archivo))
}

func leoArchivo2() {
	archivo, err := os.Open("./file-go.txt")

	if err != nil {
		// error := fmt.Errorf(err.Error())
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		// fmt.Printf("Linea > "+ registro + "\n")
		fmt.Printf("Linea > %s\n", registro)
	}

	archivo.Close()
}

func graboArchivo() {

	archivo, err := os.Create("./nuevo-archivo.txt")

	if err != nil {
		// error := fmt.Errorf(err.Error())
		fmt.Println(err)
		return
	}

	fmt.Fprintln(archivo, "Este texto se guardará en un archivo")
	archivo.Close()

}

func graboArchivo2() {

	fileName := "./nuevo-archivo.txt"

	if Append(fileName, "\nNuevos datos al archivo") == false {
		fmt.Println("erroren la segunda linea")
	}
}

func Append(fileName string, texto string) bool {
	archivo, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return false
	}

	_, errr := archivo.WriteString(texto)

	if errr != nil {
		fmt.Println(err)
		return false
	}

	return true
}
