package main

import (
	"fmt"

	"taller/conversor"
	"taller/contador"
)

func main() {
	fmt.Println("----- Bienvenido al programa -----")

	conversor.Conversor()

	fmt.Println()

	contador.Contador()
}