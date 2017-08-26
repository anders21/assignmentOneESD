package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	
	studentData := tryImportAndUnmarshal("student_data.json")

	fmt.Println(studentData)
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
