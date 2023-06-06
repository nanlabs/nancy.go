// Package utils provides the general utilities functionalities for the integrations
package utils

import (
	"bytes"
	"encoding/json"
	"strings"
)

// UnmarshalBufferToStruct unmarshals a buffer to a slice of structs with type T
func UnmarshalBufferToStruct[T any](buf bytes.Buffer) ([]T, error) {
	var outputSlice []T
	strBuf := buf.String()
	sliceStrBuf := strings.Split(strBuf, "\n")
	for _, item := range sliceStrBuf {
		if item == "" {
			continue
		}
		var m T
		err := json.Unmarshal([]byte(item), &m)
		if err != nil {
			return nil, err
		}
		outputSlice = append(outputSlice, m)
	}
	return outputSlice, nil
}
