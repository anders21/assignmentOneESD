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


func main() {
	defer profile.Start(
		profile.MemProfile,
		profile.ProfilePath("."),
	).Stop()

	fmt.Println("Sarah Anderson's Applciation")

	studentData, importErrors := tryImportAndUnmarshal("student_data_full.json")
	if importErrors != nil{
		// Crash
	}

	// Print data on the screen
	//formatted, _ := json.MarshalIndent(studentData, "", "\t")
	//fmt.Println(string(formatted))// Uncomment this line for printout of file contents
	
	valiationErrors := validateStudentData(studentData)
	if valiationErrors != nil{
		// Crash
	}

	// Generate a report, and measure the time to do so
	startTime := time.Now()
	studentReport := generateStudentReport(studentData)
	endTime := time.Now()
	

	// Print data on the screen
	formatted, _ := json.MarshalIndent(studentReport, "", "\t")
	fmt.Println(string(formatted))
	fmt.Println(startTime.Format("Mon Jan 2 2006 15:04:05.000"))
	fmt.Println(endTime.Format("Mon Jan 2 2006 15:04:05.000"))
	fmt.Print("Used time: ", endTime.Sub(startTime), "\n")
}

/*
	Validate data
	-- Check all StudentIds exist
	-- Check that important things aren't null
	-- Validate the amount of data (count( ))
*/
func validateStudentData(studentData StudentData) []error{
	var importErrors []error
	 importErrors = nil

	return importErrors
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
func tryImportAndUnmarshal(fileName string) (StudentData, []error) {
	var fileContent StudentData
	var importErrors []error
	
	// Read file and return if an error
	studentDataFileContent, readFileError := ioutil.ReadFile(fileName)
	if readFileError != nil {
		importErrors = append(importErrors, readFileError)
		return fileContent, importErrors
	}
	
	// Unmarshal data into "StudentData" struct
	jsonBlob := []byte(studentDataFileContent)
	unmarshallError := json.Unmarshal(jsonBlob, &fileContent)
	if unmarshallError != nil{
		importErrors = append(importErrors, unmarshallError)
	}
	
	return fileContent, importErrors
}

