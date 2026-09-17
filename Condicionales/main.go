package main

import "fmt"

func main() {
	flag := false
	var user string
	if flag {
		user = "Admin"
	} else {
		user = "Guest"
	}

	edad := 0

	fmt.Println("Ingresa tu edad: ")
	fmt.Scan(&edad)

	if edad >= 18 && user == "Admin" {
		fmt.Println("Hello\nTienes todos los privilegios")
	} else if edad < 18 && user == "Guest" {
		fmt.Println("Hello\nTu usuario no tiene permisos limitados")
	} else {
		fmt.Println("No tienes acceso")
	}
}
