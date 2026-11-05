package main

import "fmt"

func main() {
	for row := 1; row <= 9; row++ {
		for col := 1; col <= 9; col++ {
			fmt.Printf("%4d", row*col)
		}
		fmt.Println()
	}
}
