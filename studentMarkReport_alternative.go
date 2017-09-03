package main

import(
	"fmt"
)

/* Select all students and show the marks for each student */
func generateStudentMarkReport_alternative(studentData *StudentData) string {

	var studentReport string

	// Make a useful map
	markMap := make(map[int][]CourseMark)

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

	// Format the data ready to print
	for studentDataIndex := 0; studentDataIndex < len(studentData.MyStudents); studentDataIndex++ {
		studentReport += ("\n" + 
			studentData.MyStudents[studentDataIndex].FirstName + 
			" " + 
			studentData.MyStudents[studentDataIndex].LastName)

		for markIndex := 0; markIndex < len(markMap[studentData.MyStudents[studentDataIndex].StudentID]); markIndex++ {
			studentReport += fmt.Sprintf("\n    | %-20s|%6.2f |", 
				markMap[studentData.MyStudents[studentDataIndex].StudentID][markIndex].Class, 
				markMap[studentData.MyStudents[studentDataIndex].StudentID][markIndex].Mark)
		}
	}

	return studentReport
}