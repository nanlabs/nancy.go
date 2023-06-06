package utils

import (
	"errors"
	"reflect"
	"testing"
)

type TestCase struct {
	name    string
	in      []interface{}
	size    int
	want    [][]interface{}
	wantErr error
}

func TestBatchSlice(t *testing.T) {
	tests := []TestCase{
		{
			name: "empty slice",
			in:   []interface{}{},
			size: 1,
			want: [][]interface{}{},
		},
		{
			name: "single element slice",
			in:   []interface{}{1},
			size: 1,
			want: [][]interface{}{{1}},
		},
		{
			name: "multi-element slice",
			in:   []interface{}{1, 2, 3, 4},
			size: 2,
			want: [][]interface{}{{1, 2}, {3, 4}},
		},
		{
			name:    "invalid size",
			in:      []interface{}{1, 2, 3},
			size:    -1,
			want:    [][]interface{}{},
			wantErr: errors.New("Wrong batch size"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BatchSlice(tt.in, tt.size)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BatchSlice() = %v, want %v", got, tt.want)
			}
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("BatchSlice() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("BatchSlice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
