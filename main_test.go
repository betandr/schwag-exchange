package main

import "testing"

func TestHelloWorld(t *testing.T) {
	actualResult := "Hello"
	var expectedResult = "Hello"

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}
