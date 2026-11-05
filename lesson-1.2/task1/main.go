package main

import "fmt"

func main() {
	number := 0
	parity := ""
	isPositive := ""

	if number%2 == 0 {
		parity = "четное"
	} else {
		parity = "нечетное"
	}

	if number < 0 {
		isPositive = "отрицательное"
	} else if number > 0 {
		isPositive = "положительное"
	} else {
		isPositive = "нулевое"
	}

	fmt.Printf("Число %d: %s, %s\n", number, isPositive, parity)
}
