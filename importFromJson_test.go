package main

import(
	"testing"
	"fmt"
)
func TestReadFile(t *testing.T) {
	// Execute: Open and read content of a file
	fileContent, error := readFile("student_data_test.json")

	// Test: No errors
	AssertEqual(t, "No errors when reading contents of a file", error, nil)

	// Test: Content exists
	AssertTrue(t, "Can get content from a file", len(fileContent) > 0)
}

func TestUnmarshalJSON(t *testing.T) {
	// Set Up: Read a file
	fileContent, _ := readFile("student_data_test.json")

	// Execute: Test unmarshalJSON method
	importedStudentData, unmarshalError := unmarshalJSON(fileContent)

	// Test: No errors
	AssertEqual(t, "No errors when unmarshaling", unmarshalError, nil)

	// Test: Correct amount of data imported (Array length)
	numStudentsImported := len(importedStudentData.MyStudents)
	numMarksImported := len(importedStudentData.MyMarks)

	AssertEqual(t, "Importanted Data Student Count", numStudentsImported, 2)
	AssertEqual(t, "Importanted Data Marks Count", numMarksImported, 5)

	// Test: All imported data in correctly in the struct
	firstImportedStudent := importedStudentData.MyStudents[0]
	firstImportedStudentFirstName := firstImportedStudent.FirstName
	firstImportedStudentLastName := firstImportedStudent.LastName
	firstImportedStudentAge := firstImportedStudent.Age
	firstImportedStudentPhoneNumber := firstImportedStudent.PhoneNumber
	firstImportedStudentSuburb := firstImportedStudent.Suburb
	firstImportedStudentCity := firstImportedStudent.City

	AssertEqual(t, "Importanted Data is in correct format - First Name", firstImportedStudentFirstName, "Becka")
	AssertEqual(t, "Importanted Data is in correct format - Last Name", firstImportedStudentLastName, "Corde")

	if firstImportedStudentAge != 22 {
		t.Error("FAILED: Importanted Data is in correct format - Age")
	} else {
		fmt.Println("PASS: Importanted Data is in correct format - Age")
	}

	AssertEqual(t, "Importanted Data is in correct format - Phone Number", firstImportedStudentPhoneNumber, "026 501 3527")
	AssertEqual(t, "Importanted Data is in correct format - Suburb", firstImportedStudentSuburb, "Tamatea")
	AssertEqual(t, "Importanted Data is in correct format - City", firstImportedStudentCity, "Napier")

}

func TestValidateStudentData(t *testing.T) {
	// Set Up: Read a file and unmarshal data
	fileContent, _ := readFile("student_data_test_invalid.json")
	importedStudentData, _ := unmarshalJSON(fileContent)

	// Execute: ValidateStudentData
	validationErrors := validateStudentData(&importedStudentData)

	// Test: User is alerted to invalid data missing student
	expectedMessage := fmt.Sprintf(MissingStudentRecordValidationMessage,"Programming",0)
	actualMessage := validationErrors[0] // We only need to test the first message
	AssertEqual(t, "No student found for mark", actualMessage, expectedMessage)

	// Test: Check the count
	expectedCount := 2
	actualCount := len(validationErrors)
	AssertEqual(t, "Validation error messages count", actualCount, expectedCount)
}
