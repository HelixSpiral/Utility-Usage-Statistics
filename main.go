// Author: Shawn Smith <ShawnSmith0828@gmail.com>
// Description: Used for running some calculations on kWh usage

package main

import "fmt"

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
	var lowestDay string
	var highestDay string
	lowestDaykWh, highestDaykWh := 100.00, 0.0 // We default lowestDay to 100 and highestDay to 0
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

	fmt.Printf("Total kWh usage: %.03f\r\n", totalkWh)
	fmt.Printf("Average hourly kWh: %.03f\r\n", totalkWh/totalDataPoints)
	fmt.Printf("Average daily kWh: %.03f\r\n", totalDailykWh/totalDays)
	fmt.Printf("Lowest daily kWh: %.03f on %s\r\n", lowestDaykWh, lowestDay)
	fmt.Printf("Highest daily kWh: %.03f on %s\r\n", highestDaykWh, highestDay)
	fmt.Println("Total days:", totalDays)
	fmt.Println("Total data points:", totalDataPoints)
}
