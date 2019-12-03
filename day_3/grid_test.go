package day_3

import (
	"fmt"
	"testing"
)

func TestGetClosestIntersectionDistance(t *testing.T) {
	tests := []struct {
		d    [][]string
		want int
	}{
		{
			d: [][]string{
				{
					"U7", "R6", "D4", "L4",
				},
				{
					"R8", "U5", "L5", "D3",
				},
			},
			want: 6,
		},
		{
			d: [][]string{
				{
					"R10", "U10",
				},
				{
					"U5", "R15",
				},
			},
			want: 15,
		},
		{
			[][]string{
				{
					"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51",
				},
				{
					"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7",
				},
			},
			135,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprint("")
		t.Run(testname, func(t *testing.T) {
			actual := GetClosestIntersectionDistance(tt.d)
			if actual != tt.want {
				t.Errorf("got %d, want %d", actual, tt.want)
			}
		})
	}
}
func TestGetFewestStepsToIntersection(t *testing.T) {
	tests := []struct {
		d    [][]string
		want int
	}{
		{
			d: [][]string{
				{
					"U7", "R6", "D4", "L4",
				},
				{
					"R8", "U5", "L5", "D3",
				},
			},
			want: 30,
		},
		{
			d: [][]string{
				{
					"R10", "U10",
				},
				{
					"U5", "R15",
				},
			},
			want: 30,
		},
		{
			[][]string{
				{
					"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72",
				},
				{
					"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83",
				},
			},
			610,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprint("")
		t.Run(testname, func(t *testing.T) {
			actual := GetFewestStepsToIntersection(tt.d)
			if actual != tt.want {
				t.Errorf("got %d, want %d", actual, tt.want)
			}
		})
	}
}
