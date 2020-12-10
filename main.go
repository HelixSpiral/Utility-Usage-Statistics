// Author: Shawn Smith <ShawnSmith0828@gmail.com>
// Description: Provides some utility usage statistics to assist in alternative energy source sizing

package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/HelixSpiral/greenbuttonxml"
)

// Populated from the csv files
type meterReading struct {
	dateTime int64
	value    float64
}

func main() {
	var inputFiles []string // List of the input files we have

	// Get the current directory
	dir, _ := filepath.Abs("./")

	// Add \Input to the current path for the Input folder.
	inputFolder := fmt.Sprintf("%s\\\\Input", dir)

	// Get all the files we want to take input from
	inputFiles = returnInputFiles(inputFolder)

	meterData := getMeterData(inputFiles)

	processedData := processData(meterData)

	writeFile("processedData.txt", processedData)
}

func getMeterData(files []string) map[string]map[string]map[string][]meterReading {
	meterData := make(map[string]map[string]map[string][]meterReading)

	for _, file := range files {
		greenButtonUtilityData := greenbuttonxml.ParseGreenButtonXML(file)
		// Loop for all the readings we got
		for _, y := range greenButtonUtilityData.ServicePoint.Channel.ReadingData {

			// Parse the date and grab the year, month, and day values
			dateTime, _ := time.Parse("1/2/2006 15:04:05 PM", y.DateTime)
			yearValue := dateTime.Format("2006")
			monthValue := dateTime.Format("January")
			dayValue := dateTime.Format("02")

			// Create maps if they don't exist
			if _, ok := meterData[yearValue]; !ok {
				meterData[yearValue] = make(map[string]map[string][]meterReading)
			}
			if _, ok := meterData[yearValue][monthValue]; !ok {
				meterData[yearValue][monthValue] = make(map[string][]meterReading)
			}

			// Get the reading from the meter and append it to the slice
			meterValue, _ := strconv.ParseFloat(y.Value, 64)
			meterData[yearValue][monthValue][dayValue] = append(meterData[yearValue][monthValue][dayValue], meterReading{
				dateTime: dateTime.Unix(),
				value:    meterValue,
			})

		}

	}

	return meterData
}
