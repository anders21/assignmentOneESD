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
	StudentID float64 `json:"student_id"`
	Class     string  `json:"class"`
	Mark      float64 `json:"mark"`
}

type Student struct {
	StudentID   float64 `json:"student_id"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Age         float64 `json:"age"`
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

	startTime := time.Now()
	// Import data from JSON file, and Unmarshal into "StudentData" struct
	studentData := tryImportAndUnmarshal("student_data_test_invalid.json")
	// Generate a report
	studentReport := generateStudentReport(studentData)
	endTime := time.Now()

	// Print data on the screen
	formatted, _ := json.MarshalIndent(studentReport, "", "\t")
	fmt.Println(string(formatted))

	// Print data on the screen
	validationMessages := validateStudentData(studentData)
	if validationMessages != nil {
		fmt.Printf("Imported data with %d validation warning(s): \n", len(validationMessages))
		for messageIndex := 0; messageIndex < len(validationMessages); messageIndex++ {
			fmt.Println("* ", validationMessages[messageIndex])
		}
	}
	
	// Give a report of the time taken
	fmt.Println(startTime.Format("Mon Jan 2 2006 15:04:05.000000"))
	fmt.Println(endTime.Format("Mon Jan 2 2006 15:04:05.00000"))
	fmt.Print("Used time: ", endTime.Sub(startTime), "\n")
}

/*
	Validate data
	-- Check all marks are associated with a Student that exists
*/
func validateStudentData(studentData StudentData) []string {
	var validationMessages []string

	// Check all the ids are unique
	for marksIndex := 0; marksIndex < len(studentData.MyMarks); marksIndex ++{
		currentID := int(studentData.MyMarks[marksIndex].StudentID)
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

func studentExists(studentData StudentData, studentID int) bool{
	for studentIndex := 0; studentIndex < len(studentData.MyStudents); studentIndex ++{
		if int(studentData.MyStudents[studentIndex].StudentID) == studentID{
			return true
		}
	}
		return false
}

/*
Select all students and show the marks for each student
*/
func generateStudentReport(studentData StudentData) []StudentReportRecord {

	var studentReport []StudentReportRecord

	// Make a useful map
	markMap := make(map[float64][]CourseMark)

	for studentDataIndex := 0; studentDataIndex < len(studentData.MyMarks); studentDataIndex++ {
		currentStudentID := studentData.MyMarks[studentDataIndex].StudentID
		var currentArray []CourseMark
		currentArray = markMap[currentStudentID]
		aMark := CourseMark{
			Class: studentData.MyMarks[studentDataIndex].Class,
			Mark:  studentData.MyMarks[studentDataIndex].Mark}

		currentArray = append(currentArray, aMark)
		markMap[currentStudentID] = currentArray
	}

	// Format the data in the map
	for studentDataIndex := 0; studentDataIndex < len(studentData.MyStudents); studentDataIndex++ {
		studentReportRecord := StudentReportRecord{
			FirstName: studentData.MyStudents[studentDataIndex].FirstName,
			LastName:  studentData.MyStudents[studentDataIndex].LastName,
			Marks:     make([]CourseMark, 0)}

		for markIndex := 0; markIndex < len(markMap[studentData.MyStudents[studentDataIndex].StudentID]); markIndex++ {
			newMark := CourseMark{
				Class: markMap[studentData.MyStudents[studentDataIndex].StudentID][markIndex].Class,
				Mark:  markMap[studentData.MyStudents[studentDataIndex].StudentID][markIndex].Mark,
			}

			studentReportRecord.Marks = append(studentReportRecord.Marks, newMark)
		}
		studentReport = append(studentReport, studentReportRecord)
	}

	return studentReport
}

/*
Import data from JSON file, and save into format of "StudentData" struct
*/
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

func readFile(fileName string) (content []byte, readFileError error) {
	content, readFileError = ioutil.ReadFile(fileName)
	return
}

func unmarshallJSON(fileContent []byte) (studentData StudentData, unmarshallError error) {
	unmarshallError = json.Unmarshal(fileContent, &studentData)
	return
}
