package day_2

import (
	"reflect"
	"testing"
)

func TestParseProgram(t *testing.T) {
	data := []int {1,9,10,3,2,3,11,0,99,30,40,50}
	expected := []int {3500,9,10,70,2,3,11,0,99,30,40,50}
	actual := Parse(data)

	if !reflect.DeepEqual(expected, actual){
		t.Errorf("Slices are not equal!")
	}
}
