package main

import "fmt"

func menu() {
	var opcion string

	for {
		fmt.Println("\n----- MENÚ -----")
		fmt.Println("1. Promedio de notas de un curso")
		fmt.Println("2. Suma de números del 1 al n")
		fmt.Println("3. Convertir de Celsius a Fahrenheit")
		fmt.Println("4. Convertir de Fahrenheit a Celsius")
		fmt.Println("0 o salir. Terminar programa")
		fmt.Println("Escribe tu opción:")

		fmt.Scan(&opcion)

		if opcion == "0" || opcion == "salir" {
			fmt.Println("Saliendo del programa...")
			break
		}

		switch opcion {
		case "1":
			opcionUno()
		case "2":
			opcionDos()
		case "3":
			opcionTres()
		case "4":
			opcionCuatro()
		default:
			fmt.Println("Opción no válida, intenta de nuevo.")
		}
	}
}

func averageGrade(sumaTotal float64, cantidad int) float64 {
	return sumaTotal / float64(cantidad)
}

func opcionUno() {
	var cantidadEstudiantes int

	fmt.Println("Ingrese la cantidad de estudiantes del curso:")
	fmt.Scan(&cantidadEstudiantes)

	for cantidadEstudiantes <= 0 {
		fmt.Println("La cantidad debe ser mayor que 0. Intente de nuevo:")
		fmt.Scan(&cantidadEstudiantes)
	}

	var sumaTotal float64 = 0.0

	for i := 0; i < cantidadEstudiantes; i++ {
		var nota float64

		fmt.Printf("Ingrese la nota del estudiante %d (de 0 a 100): ", i+1)
		fmt.Scan(&nota)

		for nota < 0 || nota > 100 {
			fmt.Println("Nota inválida. Debe estar entre 0 y 100. Intente de nuevo:")
			fmt.Scan(&nota)
		}

		sumaTotal += nota
	}

	promedio := averageGrade(sumaTotal, cantidadEstudiantes)

	fmt.Printf("\nEl promedio del curso es: %.2f\n", promedio)

	if promedio >= 70 {
		fmt.Println("Estado: APROBADO")
	} else {
		fmt.Println("Estado: REPROBADO")
	}

	switch {
	case promedio >= 90:
		fmt.Println("Performance: Excellent performance")
	case promedio >= 80:
		fmt.Println("Performance: Good performance")
	case promedio >= 70:
		fmt.Println("Performance: Satisfactory performance")
	default:
		fmt.Println("Performance: Needs improvement")
	}
}

func sumarHastaN(n int) int {
	suma := 0
	for i := 1; i <= n; i++ {
		suma += i
	}
	return suma
}

func opcionDos() {
	var n int
	fmt.Println("Ingrese un número entero n:")
	fmt.Scan(&n)

	if n < 1 {
		fmt.Println("Por favor, ingrese un número mayor o igual a 1.")
		return
	}

	resultado := sumarHastaN(n)
	fmt.Printf("La suma de los números del 1 al %d es: %d\n", n, resultado)
}

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9.0 / 5.0) + 32.0
}

func opcionTres() {
	var celsius float64
	fmt.Println("Ingrese la temperatura en grados Celsius:")
	fmt.Scan(&celsius)

	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Printf("%.2f °C equivalen a %.2f °F\n", celsius, fahrenheit)
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32.0) * 5.0 / 9.0
}

func opcionCuatro() {
	var fahrenheit float64
	fmt.Println("Ingrese la temperatura en grados Fahrenheit:")
	fmt.Scan(&fahrenheit)

	celsius := fahrenheitToCelsius(fahrenheit)
	fmt.Printf("%.2f °F equivalen a %.2f °C\n", fahrenheit, celsius)
}

func main() {
	menu()
}
