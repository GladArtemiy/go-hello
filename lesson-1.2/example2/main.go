package main

import "fmt"

func main() {
	count := 0
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			count++
		}
	}
	if count > 3 {
		message := "много чётных"
		fmt.Println(message)
	} else {
		message := "мало чётных"
		fmt.Println(message)
	}
}
