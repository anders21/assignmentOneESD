package main

import (
	"fmt"
	"testing"
)

func TestGetStudentMarkList(t *testing.T) {
	// Set up: Get Student data
	fileContent, _ := readFile("student_data_testAverage.json")
	studentData, _ := unmarshalJSON(fileContent)

	// Execute: getStudentMarkList method
	locationList := getStudentMarkList(&studentData)

	// Assert: The first and the last items are correct
	AssertEqual(t, "First item in student marks: Class", locationList[0].Class, "Programming")
	AssertEqual(t, "First item in student marks: Mark", locationList[0].Mark, 10.10)
	AssertEqual(t, "First item in student marks: Location", locationList[0].Location, "Tamatea")

	AssertEqual(t, "Last item in student marks: Class", locationList[8].Class, "Operating Systems")
	AssertEqual(t, "Last item in student marks: Mark", locationList[8].Mark, 40.0) // 40.0 because Mark is of type float64
	AssertEqual(t, "Last item in student marks: Location", locationList[8].Location, "Mayfair")

	// Assert: Correct number of items
	AssertEqual(t, "Get correct number of student marks", len(locationList), 9)
}

func TestTallySuburb(t *testing.T) {
	// Set up: Get student data
	fileContent, _ := readFile("student_data_testAverage.json")
	studentData, _ := unmarshalJSON(fileContent)
	locationList := getStudentMarkList(&studentData)

	// Execute: tallySuburb method
	suburbTotalLocationList := tallySuburb(locationList)

	// Test: Tally is correct for first and last suburb
	AssertEqual(t, "Tally is correct for first suburb: Location", suburbTotalLocationList[0].Location, "Tamatea")
	AssertEqual(t, "Tally is correct for first suburb: Class", suburbTotalLocationList[0].Class, "Programming")
	AssertEqual(t, "Tally is correct for first suburb: Count", suburbTotalLocationList[0].Count, 1)
	AssertEqual(t, "Tally is correct for first suburb: Total", suburbTotalLocationList[0].Total, 10.10)

	AssertEqual(t, "Tally is correct for last suburb: Location", suburbTotalLocationList[5].Location, "Mayfair")
	AssertEqual(t, "Tally is correct for last suburb: Class", suburbTotalLocationList[5].Class, "Operating Systems")
	AssertEqual(t, "Tally is correct for last suburb: Count", suburbTotalLocationList[5].Count, 2)
	AssertEqual(t, "Tally is correct for last suburb: Total", suburbTotalLocationList[5].Total, 90.0) // 90.0 because Total is of type float64

}

func TestCalcAverage(t *testing.T) {
	// Set up: Get a list of marks
	fileContent, _ := readFile("student_data_testAverage.json")
	studentData, _ := unmarshalJSON(fileContent)
	locationList := getStudentMarkList(&studentData)
	suburbTotalLocationList := tallySuburb(locationList)

	// Execute: calcAverage method
	averages := calcAverage(suburbTotalLocationList)

	// Assert: Average is correct
	AssertEqual(t, "Programming in Napier", averages[0].Mark, 10.10)
	AssertEqual(t, "Operating Systems in Napier", averages[1].Mark, 60.0)
	AssertEqual(t, "Hardware in Napier",  averages[2].Mark, 5.50)

	AssertEqual(t, "Hardware in Hastings",  averages[3].Mark, 12.25)
	AssertEqual(t, "Programming in Hastings", fmt.Sprintf("%.4f",averages[4].Mark), fmt.Sprintf("%.4f", (11.11 + 12.12)/2)) // Had to watch the rounding error
	AssertEqual(t, "Operating Systems in Hastings",  averages[5].Mark, 45.0)
	
}


func TestSortByMarksForClassAndLocation(t *testing.T) {
	// Set up: Get a list of unordered marks
	fileContent, _ := readFile("student_data_testAverage.json")
	studentData, _ := unmarshalJSON(fileContent)
	locationList := getStudentMarkList(&studentData)
	suburbTotalLocationList := tallySuburb(locationList)
	averages := calcAverage(suburbTotalLocationList)

	// Execution: sortByMarksForClassAndLocation method
	sortedMarksForClassAndLocation := sortByMarksForClassAndLocation(averages)

	// Assert: Marks are ordered
	AssertEqual(t, "Marks are ordered correctly: First Location", sortedMarksForClassAndLocation[0].Location, "Tamatea")
	AssertEqual(t, "Marks are ordered correctly: First Class", sortedMarksForClassAndLocation[0].Class, "Operating Systems") 
	AssertEqual(t, "Marks are ordered correctly: First Mark", sortedMarksForClassAndLocation[0].Mark, 60.0) // 60.0 because Mark is of type float64

	AssertEqual(t, "Marks are ordered correctly: Last Location", sortedMarksForClassAndLocation[5].Location, "Tamatea")
	AssertEqual(t, "Marks are ordered correctly: Last Class", sortedMarksForClassAndLocation[5].Class, "Hardware") 
	AssertEqual(t, "Marks are ordered correctly: Last Mark", sortedMarksForClassAndLocation[5].Mark, 5.50)
}
