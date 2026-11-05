package main

import "fmt"

func main() {
	sum := 0
	factorial := 1
	// sumEvenNum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("Сумма всех чисел: %d\n", sum)

	sum = 0
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			continue
		}
		sum += i
	}
	fmt.Printf("Сумма четных: %d\n", sum)

	for i := 1; i <= 10; i++ {
		factorial *= i
	}
	fmt.Printf("Факториал 10: %d\n", factorial)

	sum = 0
	for i := 1; i <= 1000; i++ {
		if i%7 == 0 && i%11 == 0 {
			sum = sum + 1
		}
	}
	fmt.Printf("Делятся на 7 и 11: %d\n", sum)
}
