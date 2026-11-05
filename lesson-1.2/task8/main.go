package main

import "fmt"

func main() {
	for n := 2; n <= 100; n++ {
		if n == 2 {
			fmt.Println(n)
		} else {
			isFlag := true
			for d := 2; d < n; d++ {
				if n%d == 0 {
					isFlag = false
					break
				}
			}

			if isFlag == true {
				fmt.Println(n)
			}
		}
	}
}
