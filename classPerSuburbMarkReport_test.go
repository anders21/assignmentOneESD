package main
/*
import (
	"testing"
)


// Format the data ready to print
func TestFormatAverageMarkReport(t *testing.T) {
	// Set up: Get some data
	var locationListItems []LocationListItem

	// Execute: formatAverageMarkReport method
	formattedResult := formatAverageMarkReport(locationListItems)

	// Assert: Data is all avaliable
	AssertEqual(t, "", formattedResult, "")

}

func TestGetStudentMarkList(t *testing.T) {
	// Set up: Get Student data
	fileContent, _ := readFile("student_data_testAverage.json")
	studentData, _ := unmarshalJSON(fileContent)

	// Execute: getStudentMarkList method
	locationList := getStudentMarkList(&studentData)

	// Assert: All data is in list
	AssertEqual(t, "", locationList, "")

}

func TesttallySuburb(t *testing.T) {
	// Set up: Get student data
	fileContent, _ := readFile("student_data_testAverage.json")
	studentData, _ := unmarshalJSON(fileContent)
	locationList := getStudentMarkList(&studentData)

	// Execute: tallySuburb method
	suburbTotalLocationList := tallySuburb(locationList)

	// Test: Tally is correct
	AssertEqual(t, "", suburbTotalLocationList, "")
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
	AssertEqual(t, "Programming in Napier", averages, 10.10)
	AssertEqual(t, "Programming in Hastings", averages, 11.615)

	AssertEqual(t, "Hardware in Napier", averages, 5.50)
	AssertEqual(t, "Hardware in Hastings", averages, 12.25)

	AssertEqual(t, "Operating Systems in Napier", averages, 60)
	AssertEqual(t, "Operating Systems in Hastings", averages, 45)
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
	AssertEqual(t, "", sortedMarksForClassAndLocation, "")
}
*/