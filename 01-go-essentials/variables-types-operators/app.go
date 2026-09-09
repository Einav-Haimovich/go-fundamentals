package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount = 100
	var expoctedReturnRate float64 = 5.5
	var years int = 10
	var futureValue = float64(investmentAmount) * math.Pow(1+expoctedReturnRate/100, float64(years))
	fmt.Printf("Future Value: %.2f", futureValue)
}
