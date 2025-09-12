package main

import "fmt"

func mains() {
	var (
		nombre string
		edad   int
		pi     float64
	)

	//Entrada por teclado
	fmt.Print("Ingrese su Nombre: ")
	fmt.Scanln(&nombre) //Lee datos y guardar en la variable
	fmt.Print("Ingrese su Edad: ")
	fmt.Scanln(&edad)
	fmt.Print("Ingrese el valor de PI: ")
	fmt.Scanln(&pi)
	//Salida
	fmt.Printf("Nombre: %s Edad: %d \nValor de Pi: %f", nombre, edad, pi)

}
