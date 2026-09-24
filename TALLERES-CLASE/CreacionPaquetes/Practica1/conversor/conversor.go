package conversor

import "fmt"

func Conversor() {
	var dolares float64
	var moneda string

	fmt.Print("Ingresa la cantidad en dólares: ")
	fmt.Scan(&dolares)

	fmt.Print("Ingresa la moneda (Euros, LB, Won, BTC): ")
	fmt.Scan(&moneda)

	if moneda == "Euros" {
		fmt.Printf("Resultado: %.2f Euros\n", dolares*0.92)
	} else if moneda == "LB" {
		fmt.Printf("Resultado: %.2f Libras\n", dolares*0.79)
	} else if moneda == "Won" {
		fmt.Printf("Resultado: %.2f Wones\n", dolares*1330.0)
	} else if moneda == "BTC" {
		fmt.Printf("Resultado: %.6f BTC\n", dolares*0.000016)
	} else {
		fmt.Println("Moneda no válida.")
	}
}