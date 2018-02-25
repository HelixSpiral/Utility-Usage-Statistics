// Author: Shawn Smith <ShawnSmith0828@gmail.com>
// Description: Used for running some calculations on kWh usage
package main

import (
	"fmt"
)

type PowerData struct {
	date string
	kWh  float64
}

func main() {
	data := readData("SampleData.csv")
	var totalkWh float64
	var totalDataPoints float64

	for line := range data {
		totalkWh += data[line].kWh
		totalDataPoints += 1
	}

	fmt.Println("Total kWh usage:", totalkWh)
	fmt.Println("Total data points:", totalDataPoints)
	fmt.Println("Average kWh:", totalkWh/totalDataPoints)
}