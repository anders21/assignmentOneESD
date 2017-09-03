package main

import (
	"fmt"
	"time"

	"github.com/pkg/profile"
)

func performance() {
	defer profile.Start(
		profile.CPUProfile,
		profile.MemProfile,
		profile.ProfilePath("./Profiling"),
	).Stop()

	fmt.Println("Sarah Anderson's Application in debug mode")

	// Complete the process in two different ways
	// Import data from JSON file, and Unmarshal into "StudentData" struct and generate report
	// Only record the time to import, unmarshal and generate report, NOT print data

	// Map method
	map_startTime := time.Now()
	map_studentData := tryImportAndUnmarshal("student_data.json")
	map_studentReport := generateStudentMarkReport_alternative(&map_studentData)
	map_endTime := time.Now()

	// Array method
	array_startTime := time.Now()
	array_studentData := tryImportAndUnmarshal("student_data.json")
	array_studentReport := generateStudentMarkReport(&array_studentData)
	array_endTime := time.Now()

	// Print data on the screen
	if array_studentReport == map_studentReport {
		fmt.Println(array_studentReport)

		// Validate the data imported
		validationMessages := validateStudentData(&array_studentData)
		fmt.Println(validationMessages)		
	} else {
		fmt.Println("Reports are not the same")
		fmt.Println("Map report:")
		fmt.Println(map_studentReport)

		fmt.Println("Array report:")
		fmt.Println(array_studentReport)
	}

	// Map Report
	fmt.Println("Application used a `Mapping` method")
	fmt.Println(map_startTime.Format("Mon Jan 2 2006 15:04:05.000000"))
	fmt.Println(map_endTime.Format("Mon Jan 2 2006 15:04:05.00000"))
	fmt.Print("Used time: ", map_endTime.Sub(map_startTime), "\n")

	// Array Report
	fmt.Println("Application used an `Array` method")
	fmt.Println(array_startTime.Format("Mon Jan 2 2006 15:04:05.000000"))
	fmt.Println(array_endTime.Format("Mon Jan 2 2006 15:04:05.00000"))
	fmt.Print("Used time: ", array_endTime.Sub(array_startTime), "\n")
}
