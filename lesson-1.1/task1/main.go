package main

import (
	"fmt"
)

func main() {
	name := "Артемий"
	age := 23
	city := "Москва"
	height := 1.80
	hasJob := true

	fmt.Printf("Имя: %s\n", name)
	fmt.Printf("Возраст: %d\n", age)
	fmt.Printf("Город: %s\n", city)
	fmt.Printf("Рост: %.2f\n", height)
	fmt.Printf("Работает: %t\n", hasJob)
	fmt.Printf("Типы: %T, %T, %T, %T, %T\n", name, age, city, height, hasJob)
}
