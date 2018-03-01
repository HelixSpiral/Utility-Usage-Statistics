// Author: Shawn Smith <ShawnSmith0828@gmail.com>
// Description: Used for running some calculations on kWh usage

package main

import (
	"fmt"  // Needed to print output
	"sync" // Needed for WaitGroups
	"time" // Needed for time.Time
)

// Populated from the csv files
type PowerData struct {
	date time.Time
	kWh  float64
}

// Populated from Calculations.go
type PowerDataReturn struct {
	filePath   string // Path of the file
	lowestDay  string // Day the lowest kWh was on
	highestDay string // Day the highest kWh was on

	totalkWh         float64 // Total kWh used
	averageDailykWh  float64 // Average daily kWh used
	averageHourlykWh float64 // Average hourly kWh used
	lowestDailykWh   float64 // Lowest daily kWh found
	highestDailykWh  float64 // Highest daily kWh found
	totalDays        float64 // Total days in the file
	totalDataPoints  float64 // Total data points in the file

	hourlykWh   map[string]float64 // Hourly data
	lowestHour  map[string]float64 // Lowest kWh at each hour
	highestHour map[string]float64 // Highest kWh at each hour
}

func main() {
	var inputFiles []string // List of the input files we have
	var wg sync.WaitGroup   // Setup a waitgroup for the go routines

	var outputInfo []PowerDataReturn

	// Folder for input files
	inputFolder := "D:\\GitHub\\PowerCalculations\\Input"

	// Get all the files we want to take input from
	inputFiles = returnInputFiles(inputFolder)

	// Loop for each input file we stored
	for x := range inputFiles {
		wg.Add(1) // Add one to the waitgroup

		// Run this function inside a go routine so we can do it concurrently.
		go func(x int) {
			data := readData(inputFiles[x])
			outputInfo = append(outputInfo, runCalculations(inputFiles[x], data, &wg))
		}(x)

	}

	// Wait for the go routines to finish and then print
	wg.Wait()

	// Loop the slice of PowerDataReturns and provide output.
	for x := range outputInfo {
		fmt.Println("Power data for file:", outputInfo[x].filePath)
		fmt.Printf("Total kWh usage: %.03f\r\n", outputInfo[x].totalkWh)
		fmt.Printf("Average daily kWh: %.03f\r\n", outputInfo[x].averageDailykWh)
		fmt.Printf("Average hourly kWh: %.03f\r\n", outputInfo[x].averageHourlykWh)
		fmt.Printf("Lowest daily kWh: %.03f on %s\r\n", outputInfo[x].lowestDailykWh, outputInfo[x].lowestDay)
		fmt.Printf("Highest daily kWh: %.03f on %s\r\n", outputInfo[x].highestDailykWh, outputInfo[x].highestDay)

		// Loop for each hour and print hour-specific data.
		for y, z := range outputInfo[x].hourlykWh {
			fmt.Printf("Hour: %v | Usage: %.03f, Highest: %.03f | Lowest: %.03f\r\n", y, z/outputInfo[x].totalDays, outputInfo[x].highestHour[y], outputInfo[x].lowestHour[y])
		}
	}
}
