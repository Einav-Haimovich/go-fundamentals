package main

import (
	"fmt"
	"math"
)

func main() {
	investmentAmount, expectedReturnRate := 100, 5.5
	var years int = 10
	futureValue := float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	fmt.Printf("Future Value: %.2f", futureValue)
}
