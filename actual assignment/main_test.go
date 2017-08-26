package main

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	// Execute: Open and read content of a file
	fileContent, error := readFile("student_data_test.json")

	// Test: No errors
	AssertEqual(t, "No errors when reading contents of a file", error, nil)

	// Test: Content exists
	AssertTrue(t, "Can get content from a file", len(fileContent) > 0)
}

func TestUnmarshallJSON(t *testing.T) {
	// Set Up: Read a file
	fileContent, _ := readFile("student_data_test.json")

	// Execute: Test unmarshallJSON method
	importedStudentData, unmarshallError := unmarshallJSON(fileContent)

	// Test: No errors
	AssertEqual(t, "No errors when unmarshalling", unmarshallError, nil)

	// Test: Correct amount of data imported (Array length)
	// importedStudentData :=

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

func TestValidData(t *testing.T) {

	// Test: Correct data is imported
	// Test: User is alerted to invalid data
	//AssertEqual(t, "All fields have correct data", importedStudentData, nil)

	// Test: Correct amount of data imported (Array length)
	//AssertEqual(t, "Correct number of imported data", importedStudentData, nil)
}
