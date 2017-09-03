package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

// Validation messages
const (
	MissingStudentRecordValidationMessage = "`%s` mark cannot be imported for student with id of `%v`"
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

/* Import data from JSON file, and save into format of "StudentData" struct */
func tryImportAndUnmarshal(fileName string) StudentData {
	var fileContent StudentData

	// Read file and return if an error
	studentDataFileContent, readFileError := readFile(fileName)
	if readFileError != nil {
		panic(readFileError)
	}

	// Unmarshal data into "StudentData" struct
	fileContent, unmarshalError := unmarshalJSON(studentDataFileContent)
	if unmarshalError != nil {
		panic(unmarshalError)
	}

	return fileContent
}

/* Helper method: Read file and reutrn content and error */
func readFile(fileName string) (content []byte, readFileError error) {
	content, readFileError = ioutil.ReadFile(fileName)
	return
}

/* Helper method: Read file content, and convert from JSON to `StudentData` struct */
func unmarshalJSON(fileContent []byte) (studentData StudentData, unmarshalError error) {
	unmarshalError = json.Unmarshal(fileContent, &studentData)
	return
}

/*
	Validate StudentData imported for the following errors:
	-- Check all marks are associated with a Student that exists
*/
func validateStudentData(studentData *StudentData) []string {
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

func formatErrorMessages(validationMessages []string) string{
	report := ""
	if validationMessages != nil {
		report = fmt.Sprintf("Imported data with %d validation warning(s):", len(validationMessages))
		for messageIndex := 0; messageIndex < len(validationMessages); messageIndex++ {
			report += fmt.Sprintf("\n* %s", validationMessages[messageIndex])
		}
	} else{
		report = fmt.Sprintf("Imported data with no validation warnings.")
	}
	return report
}

/* Helper Method: Validate if an student record has been imported for a given id */
func studentExists(studentData *StudentData, studentID int) bool {
	for studentIndex := 0; studentIndex < len(studentData.MyStudents); studentIndex++ {
		if studentData.MyStudents[studentIndex].StudentID == studentID {
			return true
		}
	}
	return false
}
