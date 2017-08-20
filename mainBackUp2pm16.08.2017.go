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

	studentData := importAndUnmarshal("student_data.json")
	// Print data on the screen
	//formatted, _ := json.MarshalIndent(studentData, "", "\t")
	//fmt.Println(string(formatted))// Uncomment this line for printout of file contents

	// Validate data
	/*
	  -- Check all StudentIds exist
	  -- Check that important things aren't null
	  -- Validate the amount of data (count( ))
	*/

	//studentReport := generateStudentReport(studentData)
	//fmt.Println(studentReport)
	// Print data on the screen
	//formatted, _ := json.MarshalIndent(studentReport, "", "\t")
	//fmt.Println(string(formatted)) // Uncomment this line for printout of file contents
	generateStudentReport(studentData)

}

/*
Select all students and show the marks for each student
*/
//func generateStudentReport(studentData StudentData)([]StudentReportRecord){
func generateStudentReport(studentData StudentData) {
	//var studentReport []StudentReportRecord

	startTime := time.Now()
	fmt.Println(startTime.Format("Mon Jan 2 2006 15:04:05.000"))

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

	endTime := time.Now()
	fmt.Println(endTime.Format("Mon Jan 2 2006 15:04:05.000"))
	diff := endTime.Sub(startTime)
	fmt.Print("Used time: ", diff)
	//fmt.Println(diff.Format("Mon Jan 2 2006 15:04:05.000"))

	// Format the data in the map
	for studentDataIndex := 0; studentDataIndex < len(studentData.MyStudents); studentDataIndex++ {
		fmt.Println("\n", studentData.MyStudents[studentDataIndex].FirstName, studentData.MyStudents[studentDataIndex].LastName)

		for markIndex := 0; markIndex < len(markMap[studentData.MyStudents[studentDataIndex].StudentID]); markIndex++ {
			fmt.Printf("    | %-20s|%6.2f |\n", markMap[studentData.MyStudents[studentDataIndex].StudentID][markIndex].Class, markMap[studentData.MyStudents[studentDataIndex].StudentID][markIndex].Mark)
		}
		/*studentReportRecord := StudentReportRecord{
			FirstName: "Sarah",
			LastName: "Anderson"}
		studentReport = append(studentReport, studentReportRecord)
		*/
	}

	//return studentReport
}

/*
Import data from JSON file, and save into format of "StudentData" struct
*/
func importAndUnmarshal(fileName string) StudentData {
	// Read file and check for error
	studentDataFileContent, err := ioutil.ReadFile(fileName)
	check(err)
	// fmt.Println(string(studentDataFileContent)) // Uncomment this line for printout of file contents

	// Unmarshal data into "StudentData" struct
	jsonBlob := []byte(studentDataFileContent)
	var fileContent StudentData
	myErr := json.Unmarshal(jsonBlob, &fileContent)
	check(myErr)

	fmt.Println(fileContent)
	return fileContent
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
