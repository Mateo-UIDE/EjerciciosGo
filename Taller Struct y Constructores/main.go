package main

import (
	"fmt"
	"sort"
)

type Producto struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) Producto {
	return Producto{
		ID:       id,
		Nombre:   nombre,
		Precio:   precio,
		Cantidad: cantidad,
	}
}

func (p *Producto) ActualizarPrecio(nuevoPrecio float64) {
	if nuevoPrecio < 0 {
		fmt.Println("Error: el precio no puede ser negativo.")
		return
	}
	p.Precio = nuevoPrecio
}

func (p *Producto) ActualizarCantidad(nuevaCantidad int) {
	if nuevaCantidad < 0 {
		fmt.Println("Error: la cantidad no puede ser negativa.")
		return
	}
	p.Cantidad = nuevaCantidad
}

func (p Producto) MostrarInfo() {
	fmt.Println("----------------------------")
	fmt.Printf("ID:       %d\n", p.ID)
	fmt.Printf("Nombre:   %s\n", p.Nombre)
	fmt.Printf("Precio:   $%.2f\n", p.Precio)
	fmt.Printf("Cantidad: %d unidades\n", p.Cantidad)
	fmt.Println("----------------------------")
}

type Inventario struct {
	Productos map[int]Producto
}

func NuevoInventario() *Inventario {
	return &Inventario{
		Productos: make(map[int]Producto),
	}
}

func (inv *Inventario) AgregarProducto(p Producto) {
	if _, existe := inv.Productos[p.ID]; existe {
		fmt.Printf("Error: ya existe un producto con ID %d.\n", p.ID)
		return
	}
	inv.Productos[p.ID] = p
	fmt.Printf("Producto '%s' agregado correctamente.\n", p.Nombre)
}

func (inv *Inventario) EliminarProducto(id int) {
	if _, existe := inv.Productos[id]; !existe {
		fmt.Printf("Error: no existe un producto con ID %d.\n", id)
		return
	}
	delete(inv.Productos, id)
	fmt.Printf("Producto con ID %d eliminado.\n", id)
}

func (inv *Inventario) BuscarProducto(id int) (Producto, bool) {
	p, existe := inv.Productos[id]
	return p, existe
}

func (inv *Inventario) ActualizarProducto(p Producto) bool {
	if _, existe := inv.Productos[p.ID]; !existe {
		return false
	}
	inv.Productos[p.ID] = p
	return true
}

func (inv *Inventario) ListarProductos() {
	if len(inv.Productos) == 0 {
		fmt.Println("El inventario está vacío.")
		return
	}

	ids := make([]int, 0, len(inv.Productos))
	for id := range inv.Productos {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	fmt.Println("\n===== INVENTARIO =====")
	for _, id := range ids {
		inv.Productos[id].MostrarInfo()
	}
}


func main() {
	inventario := NuevoInventario()

	inventario.AgregarProducto(NuevoProducto(1, "Laptop", 850.50, 10))
	inventario.AgregarProducto(NuevoProducto(2, "Mouse", 15.99, 50))
	inventario.AgregarProducto(NuevoProducto(3, "Teclado", 45.00, 30))

	inventario.AgregarProducto(NuevoProducto(1, "Monitor", 200.00, 5))

	inventario.ListarProductos()

	if p, ok := inventario.BuscarProducto(2); ok {
		fmt.Println("\nProducto encontrado:")
		p.MostrarInfo()

		p.ActualizarPrecio(17.50)
		p.ActualizarCantidad(40)
		inventario.ActualizarProducto(p)
	}

	if _, ok := inventario.BuscarProducto(99); !ok {
		fmt.Println("\nEl producto con ID 99 no existe.")
	}

	inventario.EliminarProducto(3)

	fmt.Println("\nEstado final del inventario:")
	inventario.ListarProductos()
}