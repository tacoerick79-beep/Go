package main

import "fmt"

func mainfvdn() {
	var a, b int

	fmt.Print("Ingrese el primer numero: \n")
	fmt.Scanln(&a)
	fmt.Println("Ingrese el segundo numero: ")
	fmt.Scanln(&b)

	fmt.Println("¿Son Iguales? =>", a == b)
	//Distintos !=
	fmt.Println("¿Son Distintos? =>", a != b)
	//Menor que <
	fmt.Println("El Primero es menor que el segundo =>", a < b)
	//Menor o Igual que <=
	fmt.Println("El Primero es menor o igual que el segundo =>", a <= b)
	//Mayor que >
	fmt.Println("El Primero es mayor que el segundo =>", a > b)
	//Mayor o Igual que >=
	fmt.Println("El Primero es mayor o igual que el segundo =>", a >= b)
}
