package main

import "fmt"

func main() {
	expectedName := "Anna"
	exprctedAge := 30
	expectedActive := true
	actualName := "Anna"
	actualAge := 30
	actualActive := true
	passedName := expectedName == actualName
	passedAge := exprctedAge == actualAge
	passedActive := expectedActive == actualActive
	testPassed := passedName && passedAge && passedActive
	fmt.Printf("Проверка имени: %t\n", passedName)
	fmt.Printf("Проверка возраста: %t\n", passedAge)
	fmt.Printf("Проверка занятости: %t\n", passedActive)
	fmt.Printf("Тест пройден: %t\n", testPassed)
}