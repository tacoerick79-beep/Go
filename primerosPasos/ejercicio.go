package main

/*
1. Crea un sistema que pueda adivinar si es Vocal o No. Para eso una letra tiene que ser ingresada
por teclado

func main() {
	var vocal, salida string
	fmt.Println("Ingresa valor: ")
	fmt.Scanln(&vocal)

	switch vocal {
	case "a":
		salida = "Es vocal"
	case "e":
		salida = "Es vocal"
	case "i":
		salida = "Es vocal"
	case "o":
		salida = "Es vocal"
	case "u":
		salida = "Es vocal"

	default:
		salida = "No es vocal"
	}

	fmt.Println(salida)

}
*/

/*
Solicita un número e imprime todos los números pares e impares desde 1 hasta ese número con el
mensaje "es par" o "es impar" si el número es 5 el resultado será: 1 - es impar 2 - es par 3 - es
impar 4 - es par 5 - es impar

func main() {
	var numero int

	fmt.Printf("Ingrese el numero hasta donde ejecutar el ciclo de verificacion: ")
	fmt.Scanln(&numero)

	for i := 0; i <= numero; i++ {

		if i%2 == 0 {
			fmt.Println("El numero ", i, "Es par")
			continue
		} else {
			fmt.Println("El numero ", i, "Es impar")
			continue
		}
	}
}
*/

/*
3. Escriba un programa que pida un número entero mayor que cero y calcule su factorial. La factorial
es el resultado de multiplicar ese número por sus anteriores hasta la unidad.
func main() {
	var numero int
	var multiplicacion int
	fmt.Println("Ingrese un numero mayo que cero ")
	fmt.Scanln(&numero)

	for i := 0; i <= numero; i++ {
		if multiplicacion == 0 {
			multiplicacion = i
		} else {
			multiplicacion = multiplicacion * i

		}
	}

	fmt.Println(multiplicacion)
}

*/

/*
4. Escribe un programa que permita ir introduciendo una serie indeterminada de números mientras su
suma no supere 50. Cuando esto ocurra, se debe mostrar el total acumulado y el contador de
cuantos números se han introducido

func main() {
	numero := 11
	var suma, numero2, c int
	for i := 12; i > numero; i++ {
		fmt.Println("Ingrese un numero :")
		fmt.Scanln(&numero2)
		c++
		suma = suma + numero2
		if suma <= 50 {
			continue
		} else {
			fmt.Println("La contidad de numeros ingresados es: ", c)
			fmt.Println("La suma de numeros ingresados es: ", suma)
			break

		}

	}

}
*/

/*
Escribe un programa con un bucle infinito con opciones para elegir, que pueda calcular el área de 2
figuras geométricas, triángulo y rectángulo. En primer lugar, pregunta de qué figura se quiere
calcular el área, después solicita los datos que necesites para calcularlo.
1) triángulo = b * h/2
2) rectángulo = b * h
3) Salir: Solo cuando escojas esta opción el bucle se detendrá
func main() {
	inici := -1
	var b, h float32
	var opcion int
	for i := 0; i > inici; i++ {

		fmt.Println("Escojer una de las opcion 1,2,3: ")
		fmt.Println("1) triángulo = b * h/2")
		fmt.Println("2) rectángulo = b * h")
		fmt.Println("3) Salir: Solo cuando escojas esta opción el bucle se detendrá")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			fmt.Println("Ingrese la base del triangulo: ")
			fmt.Scanln(&b)
			fmt.Println("Ingrese la altura del triangulo: ")
			fmt.Scanln(&h)

			resultado := b * h / 2

			fmt.Println("El resultado es: ", resultado)
			continue

		case 2:
			fmt.Println("Ingrese la base del rectángulo: ")
			fmt.Scanln(&b)
			fmt.Println("Ingrese la altura del rectángulo: ")
			fmt.Scanln(&h)

			resultado := b * h

			fmt.Println("El resultado es: ", resultado)
			continue
		case 3:
			fmt.Println("Vuelva pronto ")
			inici = 10000000
		default:
			fmt.Println("Opcion invalida vuelva a intentar")
		}

	}
}
*/

/*
















 */
