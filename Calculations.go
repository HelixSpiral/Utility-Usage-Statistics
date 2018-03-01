package main

import "sync"

// Run the calculations on the file and return a struct of PowerDataReturn
func runCalculations(path string, data [][]PowerData, wg *sync.WaitGroup) PowerDataReturn {
	var totalkWh float64
	var totalDataPoints float64
	var totalDailykWh float64
	var totalDays float64
	var lowestDay string
	var highestDay string

	lowestDaykWh, highestDaykWh := 100.00, 0.0 // We default lowestDay to 100 and highestDay to 0
	dailykWh := make(map[string]float64)       // Track daily
	hourlykWh := make(map[string]float64)      // Track hourly
	highestHour := make(map[string]float64)    // Track highest seen per hour
	lowestHour := make(map[string]float64)     // Track lowest seen per hour

	// Loop for all the lines in data
	for _, dayData := range data {
		// Loop for each hour
		for _, hourData := range dayData {
			// Check to see if the map for that lowestHour exists, if not create it with the default value of 100
			if _, ok := lowestHour[hourData.date.Format("03:04:05 PM")]; !ok {
				lowestHour[hourData.date.Format("03:04:05 PM")] = 100
			}

			// Add the kWh usage to the totals
			totalkWh += hourData.kWh
			dailykWh[hourData.date.Format("01/02/2006")] += hourData.kWh

			// Track the totals by hour as well
			hourlykWh[hourData.date.Format("03:04:05 PM")] += hourData.kWh
			if hourData.kWh > highestHour[hourData.date.Format("03:04:05 PM")] {
				highestHour[hourData.date.Format("03:04:05 PM")] = hourData.kWh
			}

			if hourData.kWh < lowestHour[hourData.date.Format("03:04:05 PM")] && hourData.kWh > 0 {
				lowestHour[hourData.date.Format("03:04:05 PM")] = hourData.kWh
			}

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

		if dailyusage < lowestDaykWh && dailyusage > 0 {
			lowestDaykWh = dailyusage
			lowestDay = day
		}

		// Add to the total daily usage
		totalDailykWh += dailyusage

		// Increase the number of days
		totalDays += 1
	}

	wg.Done() // Tell the wg we're done.

	// Return the struct
	return PowerDataReturn{
		filePath:         path,
		lowestDay:        lowestDay,
		highestDay:       highestDay,
		totalkWh:         totalkWh,
		averageDailykWh:  totalDailykWh / totalDays,
		averageHourlykWh: totalkWh / totalDataPoints,
		lowestDailykWh:   lowestDaykWh,
		highestDailykWh:  highestDaykWh,
		totalDays:        totalDays,
		totalDataPoints:  totalDataPoints,
		hourlykWh:        hourlykWh,
		highestHour:      highestHour,
		lowestHour:       lowestHour,
	}
}
