package main

import (
	"testing"
)

func TestFindPassword(t *testing.T) {
	filename := "inputs/test_input.txt"
	password := FindPassword(filename, false, *debugMode)
	expected_password := 3

	if password != expected_password {
		t.Errorf("FindPassword(%q) = %d; want %d", filename, password, expected_password)
	}
}

func TestFindPasswordRealInput(t *testing.T) {
	filename := "inputs/input.txt"
	password := FindPassword(filename, false, *debugMode)
	expected_password := 1040

	if password != expected_password {
		t.Errorf("FindPassword(%q) = %d; want %d", filename, password, expected_password)
	}
}

func TestFindPasswordIncludeAllClicks(t *testing.T) {
	filename := "inputs/test_input.txt"
	password := FindPassword(filename, true, *debugMode)
	expected_password := 6

	if password != expected_password {
		t.Errorf("FindPassword(%q, true) = %d; want %d", filename, password, expected_password)
	}
}

func TestFindPasswordOwnTest(t *testing.T) {
	filename := "inputs/test_input2.txt"
	password := FindPassword(filename, true, *debugMode)
	expected_password := 6

	if password != expected_password {
		t.Errorf("FindPassword(%q) = %d; want %d", filename, password, expected_password)
	}
}
