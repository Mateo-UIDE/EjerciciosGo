package main

import "fmt"

func main() {
	var entero int
	var flotante float64
	fmt.Println("Ingrese dos valores: ")
	fmt.Scanf("%d %f", &entero, &flotante)

	fmt.Println("Resultado: ", (float64(entero) * flotante * flotante))
}
