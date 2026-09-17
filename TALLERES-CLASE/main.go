package main

import "fmt"

func main() {
	//EJERCICIO 1
	var n int
	fmt.Println("Ingrese un numero entero positivo:")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
	switch {
	case n >= 1 && n <= 10:
		fmt.Println("Resultado: Número pequeño")
	case n >= 11 && n <= 100:
		fmt.Println("Resultado: Número mediano")
	default:
		fmt.Println("Resultado: Número grande")
	}
	//EJERCICIO 2
	var n2 int
	fmt.Println("Ingrese un numero entero positivo:")
	fmt.Scan(&n2)
	fmt.Printf("--- Tabla de multiplicar del %d ---\n", n2)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n2, i, n2*i)
	}
	if n2%2 == 0 {
		fmt.Printf("El número %d es Par.\n", n2)
	} else {
		fmt.Printf("El número %d es Impar.\n", n2)
	}
	switch {
	case n2 >= 1 && n2 <= 5:
		fmt.Println("Clasificación: Número pequeño")
	case n2 >= 6 && n2 <= 10:
		fmt.Println("Clasificación: Número mediano")
	default:
		fmt.Println("Clasificación: Número grande")
	}
	//EJERCICIO 3
	var n3 int
	fmt.Println("Ingrese un numero entero positivo:")
	fmt.Scan(&n3)
	temp := n3
	contadorDigitos := 0
	sumaDigitos := 0

	for temp > 0 {
		digito := temp % 10
		sumaDigitos += digito
		contadorDigitos++
		temp /= 10
	}
	fmt.Printf("Cantidad de dígitos: %d\n", contadorDigitos)
	fmt.Printf("Suma de los dígitos: %d\n", sumaDigitos)
	switch contadorDigitos {
	case 1:
		fmt.Println("Mensaje: Número de una cifra")
	case 2:
		fmt.Println("Mensaje: Número de dos cifras")
	case 3:
		fmt.Println("Mensaje: Número de tres cifras")
	default:
		fmt.Println("Mensaje: Número de varias cifras")
	}
}
