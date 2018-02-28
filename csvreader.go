package main

import (
	"encoding/csv" // Needed for reading the power csv file
	"os"           // Needed for os.Open
	"strconv"      // Needed for converting kWh usage to float64
	"time"         // Needed for time parsing
)

// Read the config file and return a slice of PowerData
// [day][line]PowerData
func readData(powerFile string) [][]PowerData {
	openFile, _ := os.Open(powerFile) // Open the file
	defer openFile.Close()

	lines, err := csv.NewReader(openFile).ReadAll() // Read the file
	if err != nil {
		panic(err)
	}

	data := make([][]PowerData, 1)
	var previousDay time.Time
	var currentDay time.Time
	var dayCount int

	// Default to 1900
	previousDay, _ = time.Parse("1/2/2006 3:04:05 PM", "1/1/1900 0:00:00 PM")
	currentDay, _ = time.Parse("1/2/2006 3:04:05 PM", "1/1/1900 0:00:00 PM")

	// Loop for every line in the file
	for _, line := range lines {
		kilowatt, _ := strconv.ParseFloat(line[1], 64)             // Convert the kWh section to float64
		currentDay, _ = time.Parse("1/2/2006 3:04:05 PM", line[0]) // Convert date/time section to time.Time

		// If the current day IS the previous day continue appending data to that day
		if ok := dateCheck(previousDay, currentDay); ok != false {
			data[dayCount] = append(data[dayCount], PowerData{
				date: currentDay,
				kWh:  kilowatt,
			})
		} else {
			data = append(data, nil)
			dayCount += 1
			data[dayCount] = append(data[dayCount], PowerData{
				date: currentDay,
				kWh:  kilowatt,
			})
			previousDay = currentDay // Reset the previous day and continue
		}
	}

	return data // Return [][]PowerData
}

// Return true if event occured on the same day
func dateCheck(prev, now time.Time) bool {
	yearPrev, monthPrev, dayPrev := prev.Date()
	yearNow, monthNow, dayNow := now.Date()

	return yearPrev == yearNow && monthPrev == monthNow && dayPrev == dayNow
}
