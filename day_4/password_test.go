package day_4

import (
	"fmt"
	"testing"
)

func TestValidate(t *testing.T) {

	var tests = []struct {
		data int
		want bool
	}{
		{
			data: 122345,
			want: true,
		},
		{
			data: 111123,
			want: true,
		},
		{
			data: 1356799,
			want: true,
		},
		{
			data: 111111,
			want: true,
		}, {
			data: 223450,
			want: false,
		},
		{
			data: 123789,
			want: false,
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Data: %d", tt.data)
		t.Run(testName, func(t *testing.T) {
			actual := Validate(tt.data)

			if actual != tt.want {
				t.Errorf("got %v, want %v", actual, tt.want)
			}
		})
	}
}

func TestValidatePartTwo(t *testing.T) {
	var tests = []struct {
		data int
		want bool
	}{
		{
			data: 122345,
			want: true,
		},
		{
			data: 111123,
			want: false,
		},
		{
			data: 1356799,
			want: true,
		},
		{
			data: 111111,
			want: false,
		}, {
			data: 223450,
			want: false,
		},
		{
			data: 123789,
			want: false,
		},
		{
			data: 123444,
			want: false,
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Data: %d", tt.data)
		t.Run(testName, func(t *testing.T) {
			actual := ValidatePartTwo(tt.data)

			if actual != tt.want {
				t.Errorf("got %v, want %v", actual, tt.want)
			}
		})
	}
}
