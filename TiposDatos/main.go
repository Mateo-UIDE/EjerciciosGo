package main

import "fmt"

func main() {
	var edad int = 15
	var temperatura float64 = 21.3
	var activo bool = true
	var mensaje string = "Bienvenido"
	var datos byte = 255
	var dias [5]string = [5]string{"Lunes", "Martes", "Miercoles", "Jueves", "Viernes"}
	var numeros []float64 = []float64{1.1, 2.2, 3.3}

	fmt.Printf("Edad: %v---TipoDatos: %T\n", edad, edad)
	fmt.Printf("Temperatura: %v---TipoDatos: %T\n", temperatura, temperatura)
	fmt.Printf("Activo: %v---TipoDatos: %T\n", activo, activo)
	fmt.Printf("Mensaje: %v---TipoDatos: %T\n", mensaje, mensaje)
	fmt.Printf("Datos: %v---TipoDatos: %T\n", datos, datos)
	fmt.Printf("Dias: %v---TipoDatos: %T\n", dias, dias)
	fmt.Printf("Numeros: %v---TipoDatos: %T\n", numeros, numeros)

	fmt.Println("-------------------------------")
	fmt.Println("Tu edad es de :", edad, "Tu estado es:", activo)
	fmt.Println("La temperatura es de: ", temperatura)

}
