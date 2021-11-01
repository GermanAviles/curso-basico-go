package main

import (
	"fmt"
	"time"

	userPackage "./user"
)

type german struct {
	userPackage.Usuario
}

func main() {

	user := new(german)
	user.Create(1, "David Aviles", time.Now(), true)

	fmt.Println(user.Usuario)
}
