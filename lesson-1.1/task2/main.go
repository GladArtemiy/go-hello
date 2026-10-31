package main

import "fmt"

func main() {
	width := 5.5
	height := 3.2
	area := width * height
	perimeter := 2 * (width + height)
	fmt.Println("Ширина:", width , "Высота:", height)
	fmt.Printf("Площадь: %v\n", area)
	fmt.Printf("Периметр: %v\n", perimeter)
}
