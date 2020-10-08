package main

import "fmt"

func main() {

	var base, altura float32

	fmt.Printf("Dime cuanto mide la base del triangulo: ")
	fmt.Scanf("%f", &base)
	fmt.Printf("Dime cuanto mide la altura del triangulo: ")
	fmt.Scanf("%f", &altura)
	total := (base*altura)/2
	fmt.Println("El area del triangulo es: ", total)
}