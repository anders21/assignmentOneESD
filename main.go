package main

import (
	"fmt"
)

// When debug is "YES", the application will do profiling and timestamps
var PROFILE = "NO" // Can't `quickly` use boolean when entering a value over the command line

func main() {
	if PROFILE == "YES" {
		// Run a performance analysis on the application
		performance()

	} else {
		// Production

		fmt.Println("Sarah Anderson's Application")

		// Import data from JSON file, and Unmarshal into "StudentData" struct
		studentData := tryImportAndUnmarshal("student_data.json")

		// Validate the data imported
		validationMessages := validateStudentData(&studentData)
		fmt.Println(formatErrorMessages(validationMessages))		
		
		if len(validationMessages) == 0 {			
			// Print out marks per student
			studentReport := generateStudentMarkReport(&studentData)
			fmt.Println(studentReport)

			// Print out average mark per location per mark sorted highest to lowest
			averageReport := generateAverageClassPerSuburbMarkReport(&studentData)
			fmt.Println(averageReport)
		}
	}
}
