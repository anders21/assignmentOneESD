package main

import (
	"testing"
)

func BenchmarkMapGenerateStudentReport(b *testing.B) {
	// Set up: Get student data
	map_studentData := tryImportAndUnmarshal("student_data.json")

	// Benchmark: mapGenerateStudentReport method
	for i := 0; i < b.N; i++ {
		mapGenerateStudentReport(map_studentData)
	}
}

func BenchmarkArrayGenerateStudentReport(b *testing.B) {
	// Set up: Get student data
	array_studentData := tryImportAndUnmarshal("student_data.json")

	// Benchmark: arrayGenerateStudentReport method
	for i := 0; i < b.N; i++ {
		arrayGenerateStudentReport(array_studentData)
	}
}
