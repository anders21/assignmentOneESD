package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type StudentData struct {
	MyStudents []Student `json:"students"`
	MyMarks    []Mark    `json:"marks"`
}

type Mark struct {
	StudentID int     `json:"student_id"`
	Class     string  `json:"class"`
	Mark      float64 `json:"mark"`
}

type Student struct {
	StudentID   int    `json:"student_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Age         int    `json:"age"`
	PhoneNumber string `json:"phone_number"`
	Suburb      string `json:"suburb"`
	City        string `json:"city"`
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

		// Print out average mark per location per mark sorted highest to lowest
		averageReport := generateAverageMarkReport(studentData)
		fmt.Println(averageReport)
		
		// Print out marks per student
		studentReport := arrayGenerateStudentReport(studentData)
		fmt.Println(studentReport)

		// Validate the data imported
		validationMessages := validateStudentData(studentData)
		if validationMessages != nil {
			fmt.Printf("Imported data with %d validation warning(s): \n", len(validationMessages))
			for messageIndex := 0; messageIndex < len(validationMessages); messageIndex++ {
				fmt.Println("* ", validationMessages[messageIndex])
			}
		}
		
	}
}

/*
	Validate StudentData imported for the following errors:
	-- Check all marks are associated with a Student that exists
*/
func validateStudentData(studentData StudentData) []string {
	var validationMessages []string

	// Check all the ids are unique
	for marksIndex := 0; marksIndex < len(studentData.MyMarks); marksIndex++ {
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
func studentExists(studentData StudentData, studentID int) bool {
	for studentIndex := 0; studentIndex < len(studentData.MyStudents); studentIndex++ {
		if studentData.MyStudents[studentIndex].StudentID == studentID {
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
