// Author: Shawn Smith <ShawnSmith0828@gmail.com>
// Description: Used for running some calculations on kWh usage

package main

import (
	"fmt"           // Needed to print output
	"os"            // Needed for os.Stat
	"path/filepath" // Needed to get the current working directory
	"regexp"        // Needed for regexs
	"sync"          // Needed for WaitGroups
	"time"          // Needed for time.Time
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

	sortedKeys []string // String for the sorted hourlykWh map
}

func main() {
	var inputFiles []string // List of the input files we have
	var wg sync.WaitGroup   // Setup a waitgroup for the go routines

	var outputInfo []PowerDataReturn

	// Get the current directory
	dir, _ := filepath.Abs("./")

	// Add \Input to the current path for the Input folder.
	inputFolder := fmt.Sprintf("%s\\\\Input", dir)

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

	// Loop the slice of PowerDataReturns
	for x := range outputInfo {
		// Find the file name
		re := regexp.MustCompile("(?i)Input\\\\(.+)\\.csv")
		outputFile := re.FindStringSubmatch(outputInfo[x].filePath)

		fmt.Printf("Processing file %s...", outputFile[1])

		// Check to see if the output file already exists, if it does ignore that file and don't recalculate
		if _, err := os.Stat(fmt.Sprintf("Output\\\\%s.txt", outputFile[1])); err == nil {
			fmt.Println("Skipped - output file already found")
			continue
		}

		// Write to the file.
		err := writeFile(fmt.Sprintf("Output\\\\%s.txt", outputFile[1]), outputInfo[x])
		if err != nil {
			fmt.Println("Error:", err)
		}

		fmt.Println("Done")
	}
}
