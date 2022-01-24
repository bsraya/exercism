package main

import "fmt"

func SuccessRate(speed int) float64 {
	if speed == 0 {
		return 0.0
	} else if speed >= 1 && speed <= 4 {
		return 1.0
	} else if speed >= 5 && speed <= 8 {
		return 0.9
	} else {
		return 0.77
	}
}

func CalculateProductionRatePerHour(speed int) float64 {
	return SuccessRate(speed) * 221.0 * float64(speed)
}

func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed) / 60)
}

func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	productionRate := CalculateProductionRatePerHour(speed)
	if productionRate > limit {
		return limit
	} else {
		return productionRate
	}
}

func main() {
	ratePerHour := CalculateProductionRatePerHour(7)
	ratePerMinute := CalculateProductionRatePerMinute(5)
	fmt.Println(ratePerHour, ratePerMinute)
}
