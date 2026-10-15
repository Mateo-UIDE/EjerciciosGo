package main

import "fmt"

type Producto struct {
	Nombre string
	Precio float64
	Stock  int
}

func mostrarProducto(p Producto) {
	fmt.Println("Nombre:", p.Nombre)
	fmt.Println("Precio:", p.Precio)
	fmt.Println("Cantidad:", p.Stock)
}

func valorTotal(p Producto) float64 {
	return p.Precio * float64(p.Stock)
}

func tieneStock(p Producto) bool {
	return p.Stock > 0
}

func main() {

	p1 := Producto{Nombre: "Laptop", Precio: 850.50, Stock: 5}
	p2 := Producto{Nombre: "Mouse", Precio: 15.99, Stock: 20}
	p3 := Producto{Nombre: "Teclado", Precio: 45.00, Stock: 0}

	productos := []Producto{p1, p2, p3}

	for _, p := range productos {
		mostrarProducto(p)
		fmt.Printf("Valor total: %.2f\n", valorTotal(p))

		if tieneStock(p) {
			fmt.Println("Estado: Disponible")
		} else {
			fmt.Println("Estado: Sin stock")
		}
		fmt.Println("-------------------")
	}
}