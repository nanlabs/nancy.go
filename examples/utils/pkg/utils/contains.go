// Package utils provides the general utilities functionalities for the integrations
package utils

// Contains checks if a slice contains a specific element
func Contains[T comparable](slice []T, element T) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}
	return false
}
