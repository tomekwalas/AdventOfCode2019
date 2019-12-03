package day_3

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func GetSolution() {
	pathToFile, err := filepath.Abs("day_3/input.txt")
	data, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		panic(err)
	}
	wires := strings.Split(string(data), "\n")
	var input [][]string
	for _, wire := range wires {
		wiresPart := strings.Split(wire, ",")
		input = append(input, wiresPart)
	}

	fmt.Println(GetFewestStepsToIntersection(input))
}
