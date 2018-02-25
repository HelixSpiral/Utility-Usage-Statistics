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
	// Setup some variables
	var totalkWh float64
	var totalDataPoints float64
	var totalDailykWh float64
	var totalDays float64
	dailykWh := make(map[string]float64)

	// Read the data from the csv
	data := readData("SampleData.csv")

	// Loop for each day
	for _, day := range data {
		// Loop for each hour
		for _, hour := range day {
			// Add the kWh usage to the totals
			totalkWh += hour.kWh
			dailykWh[hour.date] += hour.kWh

			// Increase the data point total
			totalDataPoints += 1
		}
	}

	// Loop for each day
	for _, dailyusage := range dailykWh {
		// Add to the total daily usage
		totalDailykWh += dailyusage

		// Increase the number of days
		totalDays += 1
	}

	fmt.Println("Total kWh usage:", totalkWh)
	fmt.Println("Average hourly kWh:", totalkWh/totalDataPoints)
	fmt.Println("Average daily kWh:", totalDailykWh/totalDays)
	fmt.Println("Total days:", totalDays)
	fmt.Println("Total data points:", totalDataPoints)
}
