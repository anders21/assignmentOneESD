package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/pkg/profile"
)

type StudentData struct {
	MyStudents []Student `json:"students"`
	MyMarks    []Mark    `json:"marks"`
}

type Mark struct {
	StudentID int `json:"student_id"`
	Class     string  `json:"class"`
	Mark      float64 `json:"mark"`
}

type Student struct {
	StudentID   int `json:"student_id"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Age         int `json:"age"`
	PhoneNumber string  `json:"phone_number"`
	Suburb      string  `json:"suburb"`
	City        string  `json:"city"`
}

type CourseMark struct {
	Class string
	Mark  float64
}

type StudentReportRecord struct {
	FirstName string
	LastName  string
	Marks     []CourseMark
}

// Validation messages
const (
	MissingStudentRecordValidationMessage = "`%s` mark cannot be imported for student with id of `%v`"
)

func main() {
	defer profile.Start(
		profile.MemProfile,
		profile.ProfilePath("."),
	).Stop()

	fmt.Println("Sarah Anderson's Applciation")

	// Complete the process in two different ways
	// Import data from JSON file, and Unmarshal into "StudentData" struct and generate report
	// Only record the time to import, unmarshal and generate report, NOT print data

	// Map method
	map_startTime := time.Now()
	map_studentData := tryImportAndUnmarshal("student_data.json")
	map_studentReport := mapGenerateStudentReport(map_studentData)
	map_endTime := time.Now() 
	
	// Array method
	array_startTime := time.Now()
	array_studentData := tryImportAndUnmarshal("student_data.json")
	array_studentReport := arrayGenerateStudentReport(array_studentData)
	array_endTime := time.Now()
	
	// Print data on the screen
	if array_studentReport == map_studentReport	{
		fmt.Println(map_studentReport)

		// Print data on the screen
		validationMessages := validateStudentData(array_studentData)
		if validationMessages != nil {
			fmt.Printf("Imported data with %d validation warning(s): \n", len(validationMessages))
			for messageIndex := 0; messageIndex < len(validationMessages); messageIndex++ {
				fmt.Println("* ", validationMessages[messageIndex])
			}
		}
	} else{
		fmt.Println("Reports are not the same")
		fmt.Println("Map report:")
		fmt.Println(map_studentReport)

		fmt.Println("Array report:")
		fmt.Println(array_studentReport)
	}
	
	// Map Report
	fmt.Println("Applciation used a `Mapping` method")
	fmt.Println(map_startTime.Format("Mon Jan 2 2006 15:04:05.000000"))
	fmt.Println(map_endTime.Format("Mon Jan 2 2006 15:04:05.00000"))
	fmt.Print("Used time: ", map_endTime.Sub(map_startTime), "\n")

	// Array Report
	fmt.Println("Applciation used a `Array` method")
	fmt.Println(array_startTime.Format("Mon Jan 2 2006 15:04:05.000000"))
	fmt.Println(array_endTime.Format("Mon Jan 2 2006 15:04:05.00000"))
	fmt.Print("Used time: ", array_endTime.Sub(array_startTime), "\n")
}

/*
	Validate StudentData imported for the following errors:
	-- Check all marks are associated with a Student that exists
*/
func validateStudentData(studentData StudentData) []string {
	var validationMessages []string

	// Check all the ids are unique
	for marksIndex := 0; marksIndex < len(studentData.MyMarks); marksIndex ++{
		currentID := studentData.MyMarks[marksIndex].StudentID
		if !studentExists(studentData, currentID) {
			// Add an error message
			validationMessages = append(validationMessages, 
				fmt.Sprintf(
					MissingStudentRecordValidationMessage, 
					studentData.MyMarks[marksIndex].Class, 
					studentData.MyMarks[marksIndex].StudentID))
		}
	}	

	return validationMessages
}

/* Helper Method: Validate if an student record has been imported for a given id */
func studentExists(studentData StudentData, studentID int) bool{
	for studentIndex := 0; studentIndex < len(studentData.MyStudents); studentIndex ++{
		if studentData.MyStudents[studentIndex].StudentID == studentID{
			return true
		}
	}
		return false
}

/* Import data from JSON file, and save into format of "StudentData" struct */
func tryImportAndUnmarshal(fileName string) StudentData {
	var fileContent StudentData

	// Read file and return if an error
	studentDataFileContent, readFileError := readFile(fileName)
	if readFileError != nil {
		panic(readFileError)
	}

	// Unmarshal data into "StudentData" struct
	fileContent, unmarshallError := unmarshallJSON(studentDataFileContent)
	if unmarshallError != nil {
		panic(unmarshallError)
	}

	return fileContent
}

/* Helper method: Read file and reutrn content and error */
func readFile(fileName string) (content []byte, readFileError error) {
	content, readFileError = ioutil.ReadFile(fileName)
	return
}

/* Helper method: Read file content, and convert from JSON to `StudentData` struct */
func unmarshallJSON(fileContent []byte) (studentData StudentData, unmarshallError error) {
	unmarshallError = json.Unmarshal(fileContent, &studentData)
	return
}
