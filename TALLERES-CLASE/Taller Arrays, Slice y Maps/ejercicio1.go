package main

import "fmt"

func main() {

	var notas [6][4]float64

	notas = [6][4]float64{
		{8.5, 9.5, 7.5, 10.5},
		{6.4, 7.6, 8.5, 7.1},
		{9.6, 10.5, 9.8, 8.2},
		{5, 6, 7, 6},
		{8, 8, 9, 7},
		{7, 6, 8, 9},
	}

	var sumaGeneral float64

	for i := 0; i < 6; i++ {

		notasEstudiante := notas[i][:]

		suma := 0.0
		mayor := notasEstudiante[0]
		menor := notasEstudiante[0]

		for j := 0; j < 4; j++ {

			suma = suma + notasEstudiante[j]

			if notasEstudiante[j] > mayor {
				mayor = notasEstudiante[j]
			}

			if notasEstudiante[j] < menor {
				menor = notasEstudiante[j]
			}

			sumaGeneral = sumaGeneral + notasEstudiante[j]
		}

		promedio := suma / 4

		fmt.Println("Estudiante:", i+1)
		fmt.Println("Notas:", notasEstudiante)
		fmt.Println("Promedio:", promedio)
		fmt.Println("Nota mayor:", mayor)
		fmt.Println("Nota menor:", menor)
		fmt.Println()
	}

	promedioGeneral := sumaGeneral / 24

	fmt.Println("Promedio general:", promedioGeneral)
}