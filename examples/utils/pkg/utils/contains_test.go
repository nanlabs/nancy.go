// Package utils provides the general utilities functionalities for the integrations
package utils

import "testing"

// TestContainsSuccess tests the case for the slice containing the element
func TestContainsSuccess(t *testing.T) {
	slice := []string{"a", "b", "c"}
	element := "b"
	result := Contains(slice, element)
	if !result {
		t.Errorf("Expected %v to contain %v", true, result)
	}
}

// TestContainsFailure tests the case for the slice not containing the element
func TestContainsFailure(t *testing.T) {
	slice := []string{"a", "b", "c"}
	element := "z"
	result := Contains(slice, element)
	if result {
		t.Errorf("Expected %v to contain %v", false, result)
	}
}
