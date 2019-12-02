package day_2

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func firstPart(program []int) {
	result := Parse(program)

	fmt.Printf("Number is %d", result[0])
}

func secondPart(program []int) {
	totalValue := 19690720
	tempProgram := make([]int, len(program), cap(program))
	copy(tempProgram, program)
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			tempProgram[1] = i
			tempProgram[2] = j
			result := Parse(tempProgram)[0]

			if result == totalValue {
				fmt.Printf("100 * noun (%d) + verb (%d) =  %d", i, j, 100*i+j)
				return
			}
			copy(tempProgram, program)
		}

	}

	fmt.Println("Not found any pairs")
}
func GetSolution() {
	pathToFile, err := filepath.Abs("day_2/input.txt")
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
	firstPart(program)
	secondPart(program)

}
