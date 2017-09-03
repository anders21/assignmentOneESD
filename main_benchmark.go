package main

import (
	"testing"
)

func BenchmarkgenerateStudentMarkReport_alternative(b *testing.B) {
	// Set up: Get student data
	map_studentData := tryImportAndUnmarshal("student_data.json")

	// Benchmark: generateStudentMarkReport_alternative method
	for i := 0; i < b.N; i++ {
		generateStudentMarkReport_alternative(&map_studentData)
	}
}

func BenchmarkgenerateStudentMarkReport(b *testing.B) {
	// Set up: Get student data
	array_studentData := tryImportAndUnmarshal("student_data.json")

	// Benchmark: generateStudentMarkReport method
	for i := 0; i < b.N; i++ {
		generateStudentMarkReport(&array_studentData)
	}
}

