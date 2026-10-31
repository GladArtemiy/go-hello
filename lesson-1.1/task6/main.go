package main

import "fmt"

func main() {
	userName := "admin"
	password := "12345"
	attempts := float64(3)
	userName = "root"
	attempts = attempts + 1.5
	isBlocked := attempts > 3
	fmt.Printf("Пользователь: %s, пароль: %s, попыток: %.1f, блокирует: %t\n", userName, password, attempts, isBlocked)
}
