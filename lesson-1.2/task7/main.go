package main

import "fmt"

func main() {
	serverStatusOk := 200
	serverStatusBad := 503
	maxRetries := 5
	serverReadyAt := 10
	success := false

	for attempts := 1; attempts <= maxRetries; attempts++ {
		if attempts < serverReadyAt {
			fmt.Printf("Попытка %d/%d: статус %d, ждем...\n", attempts, maxRetries, serverStatusBad)
		} else {
			flag = true
			fmt.Printf("Попытка %d/%d: статус %d\n", attempts, maxRetries, serverStatusOk)
			fmt.Printf("✅ Сервер готов, попыток: %d\n", attempts)
			break
		}
	}

	if !success {
		fmt.Printf("Сервер так и не ответил после %d попыток\n", maxRetries)
	}
}
