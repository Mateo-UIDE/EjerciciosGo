package contador

import "fmt"

func Contador() {
	var frase string

	fmt.Print("Ingresa una palabra o frase sin espacios: ")
	fmt.Scan(&frase)

	a, e, i, o, u := 0, 0, 0, 0, 0

	for _, letra := range frase {
		switch letra {
		case 'a', 'A':
			a++
		case 'e', 'E':
			e++
		case 'i', 'I':
			i++
		case 'o', 'O':
			o++
		case 'u', 'U':
			u++
		}
	}

	fmt.Println("A:", a)
	fmt.Println("E:", e)
	fmt.Println("I:", i)
	fmt.Println("O:", o)
	fmt.Println("U:", u)
}