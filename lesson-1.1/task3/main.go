package main

import "fmt"

func main() {
	cost := 1500.0
	sale := 20
	saleRuble := cost * float64(sale) / float64(100)
	finalPrice := cost - saleRuble
	fmt.Printf("Цена: %.2f\n", cost)
	fmt.Printf("Скидка %d%%: %.2f\n", sale, saleRuble)
	fmt.Printf("Итого к оплате: %.2f\n", finalPrice)
}
