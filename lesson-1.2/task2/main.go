package main

import "fmt"

func main() {
	score := 50

	if score >= 90 && score <= 100 {
		fmt.Println("Оценка: отлично (5)")
	} else if score >= 75 && score <= 89 {
		fmt.Println("Оценка: хорошо (4)")
	} else if score >= 50 && score <= 74 {
		fmt.Println("Оценка: удовлетворительно (3)")
	} else if score >= 0 && score <= 49 {
		fmt.Println("Оценка: неудовлетворительно (2)")
	} else {
		fmt.Println("Ошибка: некорректный балл")
	}
}
