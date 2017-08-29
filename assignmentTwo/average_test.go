package main

import (
	"fmt"
	"testing"
)

// Format the data ready to print
func TestFormatAverageMarkReport(t *testing.T) {
	// Set up: Get some data
	someData := "fail"

	// Execute: formatAverageMarkReport method
	formattedResult := formatAverageMarkReport(someData)

	// Assert: Data is all avaliable
	AssertTrue(t, "", formattedResult)

}

func TestGetStudentMarkList(t *testing.T) {
	// Set up: Get Student data
	fileContent, _ := readFile("student_data_testAverage.json")
	studentData, _ := unmarshallJSON(fileContent)

	// Execute: getStudentMarkList method
	locationList := getStudentMarkList(studentData)

	// Assert: All data is in list
	AssertTrue(t, "", locationList)

}

func TestTellySuburb(t *testing.T) {
	// Set up: Get student data
	fileContent, _ := readFile("student_data_testAverage.json")
	studentData, _ := unmarshallJSON(fileContent)
	locationList := getStudentMarkList(studentData)

	// Execute: tellySuburb method
	suburbTotalLocationList := tellySuburb(locationList)

	// Test: Tally is correct
	AssertTrue(t, "", suburbTotalLocationList)
}

func TestCalcAverage(t *testing.T) {
	// Set up: Get a list of marks
	fileContent, _ := readFile("student_data_testAverage.json")
	studentData, _ := unmarshallJSON(fileContent)
	locationList := getStudentMarkList(studentData)
	suburbTotalLocationList := tellySuburb(locationList)

	// Execute: calcAverage method 
	averages := calcAverage(suburbTotalLocationList)

	// Assert: Average is correct
	AssertEqual(t, "Programming in Napier", averages, 10.10)
	AssertEqual(t, "Programming in Hastings", averages, 11.615)

	AssertEqual(t, "Hardware in Napier", averages, 5.50)
	AssertEqual(t, "Hardware in Hastings", averages, 12.25)

	AssertEqual(t, "Operating Systems in Napier", averages, 60)
	AssertEqual(t, "Operating Systems in Hastings", averages, 45)
)

}

func TestSortByMarksForClassAndLocation(t *testing.T) {
	// Set up: Get a list of unordered marks
	fileContent, _ := readFile("student_data_testAverage.json")
	studentData, _ := unmarshallJSON(fileContent)
	locationList := getStudentMarkList(studentData)
	suburbTotalLocationList := tellySuburb(locationList)
	averages := calcAverage(suburbTotalLocationList)

	// Execution: sortByMarksForClassAndLocation method
	sortedMarksForClassAndLocation := sortByMarksForClassAndLocation(averages)

	// Assert: Marks are ordered
	AssertTrue(t, "", sortedMarksForClassAndLocation)
}