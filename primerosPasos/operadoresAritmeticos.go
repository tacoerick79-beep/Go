package main

import "fmt"

func main() {
	var a, b float64

	//Entrada por teclado
	fmt.Print("Ingrese el numero 01: ")
	fmt.Scanln(&a)
	fmt.Print("Ingrese numero 02: ")
	fmt.Scanln(&b)

	suma := a + b
	resta := a - b
	multiplicacion := a * b
	division := a / b
	//mod := a % b

	fmt.Println("La Suma es: ", suma)
	fmt.Println("La Resta es: ", resta)
	fmt.Println("El Multiplicación es: ", multiplicacion)
	fmt.Println("La Division es: ", division)
	//fmt.Println("El mod es: ", mod)

}
