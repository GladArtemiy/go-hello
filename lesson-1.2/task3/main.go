package main

import "fmt"

func main() {
	statusCode := 211
	statusMessege := ""

	switch {
	case statusCode >= 100 && statusCode <= 199:
		statusMessege = "информационный"
		fmt.Printf("Код %d: %s", statusCode, statusMessege)
	case statusCode >= 200 && statusCode <= 299:
		statusMessege = "успех"
		fmt.Printf("Код %d: %s", statusCode, statusMessege)
	case statusCode >= 300 && statusCode <= 399:
		statusMessege = "перенаправление"
		fmt.Printf("Код %d: %s", statusCode, statusMessege)
	case statusCode >= 400 && statusCode <= 499:
		statusMessege = "ошибка клиента"
		fmt.Printf("Код %d: %s", statusCode, statusMessege)
	case statusCode >= 500 && statusCode <= 599:
		statusMessege = "ошибка сервера"
		fmt.Printf("Код %d: %s", statusCode, statusMessege)
	default:
		statusMessege = "неизвестно"
		fmt.Printf("Код %d: %s", statusCode, statusMessege)
	}

	switch statusCode {
	case 200:
		statusMessege = " - OK"
		fmt.Printf("%s", statusMessege)
	case 201:
		statusMessege = " - Created"
		fmt.Printf("%s", statusMessege)
	case 400:
		statusMessege = " - Bad Request"
		fmt.Printf("%s", statusMessege)
	case 403:
		statusMessege = " - Forbidden"
		fmt.Printf("%s", statusMessege)
	case 404:
		statusMessege = " - Not Found"
		fmt.Printf("%s", statusMessege)
	case 500:
		statusMessege = " - Internal Server Error"
		fmt.Printf("%s", statusMessege)
	}
	fmt.Println()
}
