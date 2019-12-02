package day_1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func getFuel(mass float64) float64 {
	return math.Floor(mass / float64(3)) -2
}

func GetSolution() {
	pathToFile, err := filepath.Abs("day_1/input.txt")
	check(err)

	file, err := os.Open(pathToFile)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalFuel := 0.0
	totalFuelWithFuel := 0.0
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		check(err)

		moduleFuel := getFuel(float64(mass))
		totalFuel += moduleFuel

		for moduleFuel > 0 {
			totalFuelWithFuel += moduleFuel
			moduleFuel = getFuel(moduleFuel)
		}

	}
	fmt.Printf("Total fuel: %.2f\n", totalFuel)
	fmt.Printf("Total fuel with fuel: %.2f", totalFuelWithFuel)
}