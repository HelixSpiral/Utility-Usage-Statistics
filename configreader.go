package main

import (
	"encoding/csv" // Needed for reading the power csv file
	"os"
	"strconv"
)

// Read the config file and return a slice of PowerData
func readData(powerFile string) []PowerData {
	openFile, _ := os.Open(powerFile) // Open the file
	defer openFile.Close()

	lines, err := csv.NewReader(openFile).ReadAll() // Read the file
	if err != nil {
		panic(err)
	}

	var data []PowerData

	for _, line := range lines {
		kilowatt, _ := strconv.ParseFloat(line[1], 64) // Convert the kWh section to float64
		data = append(data, PowerData{
			date: line[0],
			kWh:  kilowatt,
		})
	}

	return data // Return []PowerData
}