package main

import (
	"fmt"
	"sort"
)

type LocationListItem struct{
	Location string
	Mark float64
	Class string
}

type LocationClassGroup struct{
	Location string
	Class string
	Count int 
	Total float64
}
/*
A function that provides an overview (in text format) of the average mark per paper
per suburb, sorted on marks from high to low.
*/
func generateAverageClassPerSuburbMarkReport(studentData *StudentData) string {
	locationList := getStudentMarkList(studentData)
	suburbTotalLocationList := tallySuburb(locationList)
	averages := calcAverage(suburbTotalLocationList)
	sortedMarksForClassAndLocation := sortByMarksForClassAndLocation(averages)
	return formatAverageMarkReport(sortedMarksForClassAndLocation)
}

// Format the data ready to print
func formatAverageMarkReport(locationList []LocationListItem) string {
	var studentReport string
	studentReport += fmt.Sprintf("\n\t| %-20s | %-20s| %-12s |", 
		"Location", 
		"Class", 
		"Average Mark")

	for markIndex := 0; markIndex < len(locationList); markIndex++ {
		studentReport += fmt.Sprintf("\n\t| %-20s | %-20s| %12.2f |", 
			locationList[markIndex].Location, 
			locationList[markIndex].Class, 
			locationList[markIndex].Mark)
	}
	
	return studentReport
}

func getStudentMarkList(studentData *StudentData) []LocationListItem{
	
	// Imitate the map using a look up table for the studentID and an array of an array of marks
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

	var dtoList []LocationListItem

	// Format the data ready to print
	for studentDataIndex := 0; studentDataIndex < len(studentData.MyStudents); studentDataIndex++ {	

			studentIndex := getIndexForStudentID(lookUpTable, studentData.MyStudents[studentDataIndex].StudentID)			
			studentMarkArray := markArray[studentIndex]

		for markIndex := 0; markIndex < len(studentMarkArray); markIndex++ {
			
			dtoItem := LocationListItem{
				Location: studentData.MyStudents[studentDataIndex].Suburb,
				Class: studentMarkArray[markIndex].Class,
				Mark:  studentMarkArray[markIndex].Mark}

			dtoList = append(dtoList, dtoItem)
		}
	}

	return dtoList
}

func tallySuburb(locationList []LocationListItem) []LocationClassGroup{
	var groupList []LocationClassGroup

	for index := 0; index < len(locationList); index ++ {
		location := locationList[index].Location
		class := locationList[index].Class
		mark := locationList[index].Mark
		groupID := getLocationClassGroupID(groupList, location, class)

		if groupID != -1 {
			groupList[groupID].Count += 1
			groupList[groupID].Total += mark

		}else{
			groupItem := LocationClassGroup{
				Location: location,
				Class: class,
				Count: 1,
				Total: mark}
			groupList = append(groupList, groupItem)
		}
	}
	return groupList	
}

func calcAverage(groupList []LocationClassGroup) []LocationListItem{
	var records []LocationListItem

	for index := 0; index < len(groupList); index ++ {
		newRecord := LocationListItem{
			Location: groupList[index].Location,
			Class: groupList[index].Class,
			Mark: (groupList[index].Total/float64(groupList[index].Count))} // Average = `Total` divided by `Count`
		
		records = append(records, newRecord)
	}

	return records
}

func sortByMarksForClassAndLocation(records []LocationListItem) []LocationListItem{

	sort.Slice(records, func(i, j int) bool { 
		return records[i].Mark > records[j].Mark })

	return records
}

func getLocationClassGroupID(groupList []LocationClassGroup, location string, class string) int{
	for index := 0; index < len(groupList); index ++ {
		if groupList[index].Location == location && groupList[index].Class == class{
			return index
		}
	}
	return -1
}