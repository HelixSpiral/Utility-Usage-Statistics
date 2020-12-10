package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
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

		// Make sure the file is a .xml
		if strings.Contains(path, ".xml") {
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
func writeFile(file string, data UtilityDataStatistics) error {
	var currentMonth string
	// Create the file
	createdFile, err := os.Create(file)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer createdFile.Close()

	// Create a new writer for the file
	write := bufio.NewWriter(createdFile)

	// Write the data to the file
	write.WriteString(fmt.Sprintf("Processed %d hours in %d days\r\n", int(data.totalDataPoints), int(data.totalDays)))
	write.WriteString(fmt.Sprintf("Total usage: %.03f kWh\r\n", data.totalkWh))
	write.WriteString(fmt.Sprintf("Lowest daily usage: %.03f kWh\r\n", data.lowestDailykWh))
	write.WriteString(fmt.Sprintf("Highest daily usage: %.03f kWh\r\n", data.highestDailykWh))
	write.WriteString("----------\r\n")

	sortedKeys := sortKeys(data.dailykWh)

	for _, y := range sortedKeys {
		timeParse, _ := time.Parse("2006/01/02", y)
		if timeParse.Format("January") != currentMonth {
			write.WriteString(fmt.Sprintf("Summary - %s\r\n", timeParse.Format("2006/January")))
			currentMonth = timeParse.Format("January")
		}
		write.WriteString(fmt.Sprintf("%s: Total: %.03f kWh | Lowest Hour: %.03f | Highest Hour: %.03f\r\n", y, data.dailykWh[y], data.lowestHour[y], data.highestHour[y]))
	}

	// Flush
	write.Flush()

	// Return no error
	return nil
}
