package main

import "fmt"

func ganador(votos map[string]int) string {

	actividad := ""
	mayor := 0

	for nombre, cantidad := range votos {

		if cantidad > mayor {
			mayor = cantidad
			actividad = nombre
		}
	}

	return actividad
}

func main() {

	votos := map[string]int{
		"deportes":    0,
		"videojuegos": 0,
		"cine":        0,
		"musica":      0,
	}

	for i := 0; i < 5; i++ {

		var voto string

		fmt.Println("Actividades: deportes, videojuegos, cine, musica")
		fmt.Print("Ingrese su voto: ")
		fmt.Scan(&voto)

		if voto == "deportes" {
			votos["deportes"]++
		} else if voto == "videojuegos" {
			votos["videojuegos"]++
		} else if voto == "cine" {
			votos["cine"]++
		} else if voto == "musica" {
			votos["musica"]++
		} else {
			fmt.Println("Voto no valido")
			i--
		}
	}

	fmt.Println()
	fmt.Println("Resultados:")

	for actividad, cantidad := range votos {
		fmt.Println(actividad, ":", cantidad)
	}

	resultado := ganador(votos)

	fmt.Println()
	fmt.Println("Actividad ganadora:", resultado)
	fmt.Println("Votos:", votos[resultado])
}