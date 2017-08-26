package main

import (
	"fmt"
	"testing"
)

/*Our assert library*/
func AssertEqual(t *testing.T, message string, item1, item2 interface{}) {
	if item1 != item2 {
		t.Error("FAILED:", message, "- item1:", item1, "item2:", item2)
	} else {
		fmt.Println("PASS:", message)
	}
}
func AssertNotEqual(t *testing.T, message string, item1, item2 interface{}) {
	if item1 == item2 {
		t.Error("FAILED:", message, "- item1:", item1, "item2:", item2)
	} else {
		fmt.Println("PASS:", message)
	}
}

func AssertTrue(t *testing.T, message string, item bool) {
	if item {
		fmt.Println("PASS:", message)
	} else {
		t.Error("FAILED:", message, "- item:", item)
	}
}

func AssertFalse(t *testing.T, message string, item bool) {
	AssertTrue(t, message, !item)
}
