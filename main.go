package main

import "fmt"

func main() {
	var base, altura float32
	fmt.Scanf("%f", &base)
	fmt.Scanf("%f", &altura)
	area := (base*altura)/2
	fmt.Println(area)
}