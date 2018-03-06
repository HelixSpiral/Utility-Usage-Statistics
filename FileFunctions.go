package main

import (
	"bufio"         // Needed to write to the output file
	"fmt"           // Used for printing error messages
	"os"            // Needed for os.FileInfo
	"path/filepath" // Needed for filepath.Walk
	"strings"       // Needed for strings.Contains
)

// Get a list of the files we want to take input from
func returnInputFiles(inputFolder string) []string {
	var inputFiles []string
	err := filepath.Walk(inputFolder, func(path string, fileInfo os.FileInfo, err error) error {

		// Check for any errors
		if err != nil {
			fmt.Println("Error accessing:", err)
			return err
		}

		// Ignore any folder not our input folder - we're not doing recursive searching
		if fileInfo.IsDir() && fileInfo.Name() != "Input" {
			return filepath.SkipDir
		}

		// Make sure the file is a .csv
		if strings.Contains(path, ".csv") {
			inputFiles = append(inputFiles, path)
		}

		// Return no error if we made it this far
		return nil
	})

	// Error handling for the above function
	if err != nil {
		fmt.Println("Error:", err)
	}

	return inputFiles
}

// Write to the file
func writeFile(file string, data PowerDataReturn) error {
	timePM := false
	// Create the file
	createdFile, err := os.Create(file)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer createdFile.Close()

	// Create a new writer for the file
	write := bufio.NewWriter(createdFile)

	// Write the data to the file
	write.WriteString(fmt.Sprintf("Power data for file: %s\r\n", data.filePath))
	write.WriteString(fmt.Sprintf("Total kWh usage: %.03f\r\n", data.totalkWh))
	write.WriteString(fmt.Sprintf("Average daily kWh: %.03f\r\n", data.averageDailykWh))
	write.WriteString(fmt.Sprintf("Average hourly kWh: %.03f\r\n", data.averageHourlykWh))
	write.WriteString(fmt.Sprintf("Lowest daily kWh: %.03f on %s\r\n", data.lowestDailykWh, data.lowestDay))
	write.WriteString(fmt.Sprintf("Highest daily kWh: %.03f on %s\r\n", data.highestDailykWh, data.highestDay))

	// Bit of a hack to print AM before PM
Loop:
	for _, y := range data.sortedKeys {
		if strings.Contains(y, "AM") && timePM != true {
			write.WriteString(fmt.Sprintf("Hour: %v | Usage: %.03f, Highest: %.03f | Lowest: %.03f\r\n", y, data.hourlykWh[y]/data.totalDays, data.highestHour[y], data.lowestHour[y]))
			if y == "12:00:00 AM" {
				timePM = true
				goto Loop
			}
		} else if strings.Contains(y, "PM") && timePM == true {
			write.WriteString(fmt.Sprintf("Hour: %v | Usage: %.03f, Highest: %.03f | Lowest: %.03f\r\n", y, data.hourlykWh[y]/data.totalDays, data.highestHour[y], data.lowestHour[y]))
		}
	}

	// Flush
	write.Flush()

	// Return no error
	return nil
}
