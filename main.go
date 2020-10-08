package main

import "fmt"

func main() {
	var base, altura int
	fmt.Scanf("%d", &base)
	fmt.Scanf("%d", &altura)
	area := (base*altura)/2
	fmt.Println(area)
}