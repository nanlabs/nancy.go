// Package utils provides the general utilities functionalities for the integrations
package utils

import "errors"

// BatchSlice: Splits a slice into batches of a given size.
func BatchSlice[T any](in []T, size int) ([][]T, error) {
	out := make([][]T, 0)

	if size <= 0 {
		return out, errors.New("Wrong batch size")
	}

	for i := 0; i < len(in); i = i + size {
		j := i + size
		if j > len(in) {
			j = len(in)
		}
		out = append(out, in[i:j])
	}

	return out, nil
}
