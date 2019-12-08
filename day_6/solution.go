package day_6

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func GetSolution() {
	pathToFile, err := filepath.Abs("day_6/input.txt")
	data, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		panic(err)
	}
	var orbits []Orbit
	for _, orbitString := range strings.Split(string(data), "\n") {
		orbitString = strings.Replace(orbitString, "\r", "", -1)
		s := strings.Split(orbitString, ")")
		orbits = append(orbits, Orbit{
			center: s[0],
			child:  s[1],
		})
	}

	orbitsCount := GetCountOfOrbits(orbits)
	orbitTransfers := GetOrbitalTransfers(orbits)

	fmt.Println(orbitsCount)
	fmt.Println(orbitTransfers)
}
