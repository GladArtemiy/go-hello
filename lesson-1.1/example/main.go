package main

import (
	"fmt"
	"strconv"
)

// Константа: адрес тестируемого сервера, он не меняется
const BaseURL = "https://api.shop.com"

func main() {
	// --- Данные о товаре ---
	productName := "Наушники"
	price := 2499.90
	quantity := 3
	inStock := true

	// --- Расчёт ---
	total := price * float64(quantity) // int → float64, иначе Go не даст перемножить

	fmt.Printf("Товар: %s\n", productName)
	fmt.Printf("Цена: %.2f руб., количество: %d\n", price, quantity)
	fmt.Printf("Итого: %.2f руб.\n", total)
	fmt.Printf("В наличии: %t\n", inStock)

	// --- Типы переменных ---
	fmt.Printf("Типы: %T, %T, %T, %T\n", productName, price, quantity, inStock)

	// --- Мини-автотест ---
	productID := 15
	url := BaseURL + "/products/" + strconv.Itoa(productID)
	fmt.Println("Запрос на:", url)

	expectedStatus := 200
	actualStatus := 200 // пока «притворяемся», что сервер ответил 200
	passed := expectedStatus == actualStatus

	fmt.Printf("Ожидали %d, получили %d → тест пройден: %t\n",
		expectedStatus, actualStatus, passed)
}