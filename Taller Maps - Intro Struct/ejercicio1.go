package main

import "fmt"

func main() {
	conteo := make(map[string]int)

	fmt.Println("Ingrese su texto y escriba 'fin' al terminar:")

	for {
		var palabra string
		_, err := fmt.Scan(&palabra)
		if err != nil || palabra == "fin" {
			break
		}
		conteo[palabra]++
	}

	fmt.Println("\nConteo de palabras:")
	for palabra, veces := range conteo {
		fmt.Printf("%s: %d\n", palabra, veces)
	}
	
	masRepetida := ""
	maximo := 0
	for palabra, veces := range conteo {
		if veces > maximo {
			maximo = veces
			masRepetida = palabra
		}
	}

	fmt.Printf("\nLa palabra que más se repite es '%s' con %d repeticiones\n", masRepetida, maximo)
}

