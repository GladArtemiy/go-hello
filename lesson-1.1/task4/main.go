package main

import "fmt"

func main() {
	seconds := 5001
	hours := seconds / 3600
	hoursInSeconds := hours * 3600
	minutes := (seconds - hoursInSeconds) / 60
	minutesInSeconds := (minutes * 60)
	finalSeconds := seconds - hoursInSeconds - minutesInSeconds
	fmt.Printf("%d Секунд = %dч %dм %dс", seconds, hours, minutes, finalSeconds)
}
