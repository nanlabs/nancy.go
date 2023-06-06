package utils

import (
	"bytes"
	"encoding/json"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Test case 1: Empty buffer
func TestUnmarshalBufferToStructEmptyBuffer(t *testing.T) {
	emptyBuf := bytes.Buffer{}
	emptySlice, err := UnmarshalBufferToStruct[Person](emptyBuf)
	if err != nil {
		t.Errorf("UnmarshalBufferToStruct() returned an error for empty buffer: %v", err)
	}
	if len(emptySlice) != 0 {
		t.Errorf("UnmarshalBufferToStruct() returned non-empty slice for empty buffer: %v", emptySlice)
	}
}

// Test case 2: Buffer with one valid JSON object
func TestUnmarshalBufferToStructWithOneValidJsonObj(t *testing.T) {
	buf := bytes.Buffer{}
	p := Person{Name: "Alice", Age: 30}
	err := json.NewEncoder(&buf).Encode(p)
	if err != nil {
		t.Fatalf("Error encoding Person: %v", err)
	}
	singleSlice, err := UnmarshalBufferToStruct[Person](buf)
	if err != nil {
		t.Errorf("UnmarshalBufferToStruct() returned an error for single JSON object buffer: %v", err)
	}
	if len(singleSlice) != 1 {
		t.Errorf("UnmarshalBufferToStruct() returned incorrect number of objects for single JSON object buffer: %v", singleSlice)
	}
	if singleSlice[0].Name != "Alice" || singleSlice[0].Age != 30 {
		t.Errorf("UnmarshalBufferToStruct() returned incorrect object for single JSON object buffer: %v", singleSlice[0])
	}
}

// Test case 3: Buffer with multiple valid JSON objects
func TestUnmarshalBufferToStructWithMultipleValidJsonObjs(t *testing.T) {
	buf := bytes.Buffer{}
	p1 := Person{Name: "Alice", Age: 30}
	p2 := Person{Name: "Bob", Age: 40}
	err := json.NewEncoder(&buf).Encode(p1)
	if err != nil {
		t.Fatalf("Error encoding Person 1: %v", err)
	}
	err = json.NewEncoder(&buf).Encode(p2)
	if err != nil {
		t.Fatalf("Error encoding Person 2: %v", err)
	}
	multiSlice, err := UnmarshalBufferToStruct[Person](buf)
	if err != nil {
		t.Errorf("UnmarshalBufferToStruct() returned an error for multiple JSON objects buffer: %v", err)
	}
	if len(multiSlice) != 2 {
		t.Errorf("UnmarshalBufferToStruct() returned incorrect number of objects for multiple JSON objects buffer: %v", multiSlice)
	}
	if multiSlice[0].Name != "Alice" || multiSlice[0].Age != 30 {
		t.Errorf("UnmarshalBufferToStruct() returned incorrect object for multiple JSON objects buffer: %v", multiSlice[0])
	}
	if multiSlice[1].Name != "Bob" || multiSlice[1].Age != 40 {
		t.Errorf("UnmarshalBufferToStruct() returned incorrect object for multiple JSON objects buffer: %v", multiSlice[1])
	}
}

// Test case 4: Buffer with invalid JSON object
func TestUnmarshalBufferToStructWithInvalidJsonObj(t *testing.T) {
	buf := bytes.Buffer{}
	invalidStr := "invalid JSON object"
	_, err := buf.WriteString(invalidStr)
	if err != nil {
		t.Fatalf("Error writing invalid JSON object: %v", err)
	}
	invalidSlice, err := UnmarshalBufferToStruct[Person](buf)
	if err == nil {
		t.Errorf("UnmarshalBufferToStruct() did not return an error for invalid JSON object buffer: %v", invalidSlice)
	}
}
