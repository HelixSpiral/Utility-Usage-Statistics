// Author: Shawn Smith <ShawnSmith0828@gmail.com>
// Description: Used for running some calculations on kWh usage
package main

import (
	"fmt"
	"time"
)

type PowerData struct {
	date time.Time
	kWh  float64
}

func main() {
	data := readData("SampleData.csv")
	var totalkWh float64
	var totalDataPoints float64

	for _, day := range data {
		for _, hour := range day {
			totalkWh += hour.kWh
			totalDataPoints += 1
		}
	}

	fmt.Println("Total kWh usage:", totalkWh)
	fmt.Println("Total data points:", totalDataPoints)
	fmt.Println("Average kWh:", totalkWh/totalDataPoints)
}
