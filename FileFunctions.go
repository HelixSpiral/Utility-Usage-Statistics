package main

import (
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
