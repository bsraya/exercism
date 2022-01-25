package main

import "fmt"

func leapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	result := leapYear(2100)
	fmt.Println(result)
}
