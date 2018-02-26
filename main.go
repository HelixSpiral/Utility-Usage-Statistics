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
	// Setup some variables
	var totalkWh float64
	var totalDataPoints float64
	var totalDailykWh float64
	var totalDays float64
	var lowestDay string
	var highestDay string
	lowestDaykWh, highestDaykWh := 100.00, 0.0 // We default lowestDay to 100 and highestDay to 0
	dailykWh := make(map[string]float64)
	hourlykWh := make([]float64, 24)

	// Read the data from the csv
	data := readData("PowerData.csv")

	// Loop for each day
	for _, dayData := range data {
		// Loop for each hour
		for hour, hourData := range dayData {
			// Add the kWh usage to the totals
			totalkWh += hourData.kWh
			dailykWh[hourData.date.Format("01/02/2006")] += hourData.kWh

			// Track the totals by hour as well
			hourlykWh[hour] += hourData.kWh

			// Increase the data point total
			totalDataPoints += 1
		}
	}

	// Loop for each day
	for day, dailyusage := range dailykWh {
		if dailyusage > highestDaykWh {
			highestDaykWh = dailyusage
			highestDay = day
		}

		if dailyusage < lowestDaykWh {
			lowestDaykWh = dailyusage
			lowestDay = day
		}

		// Add to the total daily usage
		totalDailykWh += dailyusage

		// Increase the number of days
		totalDays += 1
	}

	// Print the data
	fmt.Printf("Total kWh usage: %.03f\r\n", totalkWh)
	fmt.Printf("Average hourly kWh: %.03f\r\n", totalkWh/totalDataPoints)
	fmt.Printf("Average daily kWh: %.03f\r\n", totalDailykWh/totalDays)
	fmt.Printf("Lowest daily kWh: %.03f on %s\r\n", lowestDaykWh, lowestDay)
	fmt.Printf("Highest daily kWh: %.03f on %s\r\n", highestDaykWh, highestDay)
	fmt.Println("Total days:", totalDays)
	fmt.Println("Total data points:", totalDataPoints)

	// Loop for each hour and print the average kWh usage for that hour
	for x, y := range hourlykWh {
		fmt.Printf("Hour: %02d | Usage: %.03f\r\n", x, y/totalDays)
	}
}
