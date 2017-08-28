package main

import (
)

type LocationListItem struct{
	Location string
	Mark float64
	Class string
}

// Ability to hold student and marks together
type StudentMarksDTO struct{
	StudentID int
	Suburb string
	Marks []MarkRecord	
}

type MarkRecord struct {
	Class string  
	Mark float64 
}

/*
A function that provides an overview (in text format) of the average mark per paper
per suburb, sorted on marks from high to low.
*/
func generateAverageMarkReport(studentMarksList []StudentMarksDTO) string {

}


// We take a copy of the studentData so that we can delete the marks we have already processed
func getStudentMarkList(studentData StudentData) []StudentMarksDTO{
	var studentMarksDTOList []StudentMarksDTO

	// For each student with marks, collect their marks
	for studentDataIndex := 0; studentDataIndex < len(studentData.MyMarks); studentDataIndex++ {
		currentStudentID := studentData.MyMarks[studentDataIndex].StudentID
		currentStudentMarkDTO := getStudentMarkDTO(studentMarksDTOList, currentStudentID)

		aMark := MarkRecord{
			Class: studentData.MyMarks[studentDataIndex].Class,
			Mark:  studentData.MyMarks[studentDataIndex].Mark}

		var studentMarkRecords []MarkRecord
		studentMarkRecords = append(studentMarkRecords, currentArray)
		currentArray = studentMarkRecords[arrayIndex]
		aMark := CourseMark{
			Class: studentData.MyMarks[studentDataIndex].Class,
			Mark:  studentData.MyMarks[studentDataIndex].Mark}

		studentMarkRecords[arrayIndex] = append(currentArray, aMark)
	}



	// Imitate the map using a look up table for the studentID and an array of an array of marks
	var lookUpTable []int
	var markArray [][]CourseMark

	for studentDataIndex := 0; studentDataIndex < len(studentData.MyMarks); studentDataIndex++ {
		
		
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
}

// Get the pointer to the studentDTO record		
func getStudentMarkDTO(studentMarksDTOList []StudentMarksDTO, studentID int) StudentMarksDTO{

}