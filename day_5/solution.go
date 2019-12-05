package day_5

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func GetSolution() {
	pathToFile, err := filepath.Abs("day_5/input.txt")
	data, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		panic(err)
	}
	programString := strings.Split(string(data), ",")
	var program []int
	for _, inputString := range programString {
		input, err := strconv.Atoi(inputString)
		if err != nil {
			panic(err)
		}
		program = append(program, input)
	}

	RunPrograms(program, 5)
}
