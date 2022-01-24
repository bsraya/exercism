package main

import "fmt"

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

func CalculateCost(carsCount int) uint {
	var cost uint
	if carsCount < 10 {
		cost = uint(carsCount * 10000)
	} else {
		divisibleByTen := carsCount / 10
		rest := carsCount % (divisibleByTen * 10)
		cost = uint((divisibleByTen * 95000) + (rest * 10000))
	}
	return cost
}

func main() {
	ratePerHour := CalculateWorkingCarsPerHour(1547, 90)
	ratePerMinute := CalculateWorkingCarsPerMinute(1105, 90)
	cost := CalculateCost(37)
	fmt.Println(ratePerHour, ratePerMinute, cost)
}
