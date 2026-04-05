package main

import "fmt"

func main() {
	const pi = 3.14
	var radius float64 = 5.0

	area := pi * radius * radius
	fmt.Printf("半径为 %.1f的圆，面积是 %.2f", radius, area)
}
