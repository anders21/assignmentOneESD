package main

import(
	"fmt"
)


/* Select all students and show the marks for each student */
func generateStudentMarkReport(studentData *StudentData) string {

	var studentReport string

	// Immidate the map using a look up table for the studentID and an array of an array of marks
	var lookUpTable []int
	var markArray [][]CourseMark

	for studentDataIndex := 0; studentDataIndex < len(studentData.MyMarks); studentDataIndex++ {
		currentStudentID := studentData.MyMarks[studentDataIndex].StudentID
		lookUpTable = append(lookUpTable, currentStudentID)
		arrayIndex := getIndexForStudentID(lookUpTable, currentStudentID)
	
		var currentArray []CourseMark
		markArray = append(markArray, currentArray)
		currentArray = markArray[arrayIndex]
		aMark := CourseMark{
			Class: studentData.MyMarks[studentDataIndex].Class,
			Mark:  studentData.MyMarks[studentDataIndex].Mark}

		markArray[arrayIndex] = append(currentArray, aMark)
	}

	// Format the data ready to print
	for studentDataIndex := 0; studentDataIndex < len(studentData.MyStudents); studentDataIndex++ {
		studentIndex := getIndexForStudentID(lookUpTable, studentData.MyStudents[studentDataIndex].StudentID)			
		studentMarkArray := markArray[studentIndex]

		for markIndex := 0; markIndex < len(studentMarkArray); markIndex++ {
			studentReport += fmt.Sprintf("\n| %-20s| %-20s| %-20s|%6.2f |", 
				studentData.MyStudents[studentDataIndex].FirstName,
				studentData.MyStudents[studentDataIndex].LastName,
				studentMarkArray[markIndex].Class, 
				studentMarkArray[markIndex].Mark)
		}
	}

	return studentReport
}

// Takes a studentID and trys to find it's location in the lookUpTable which relates to the Id of the mark array
func getIndexForStudentID(lookUpTable []int, studentID int) int {
	for index := 0; index < len(lookUpTable); index ++ {
		if lookUpTable[index] == studentID {
			return index
		}
	}
	return 1
}