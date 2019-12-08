package day_6

import (
	"strings"
	"testing"
)

func prepareTestData(data []string) []Orbit {
	var orbits []Orbit
	for _, orbitString := range data {
		s := strings.Split(orbitString, ")")
		orbits = append(orbits, Orbit{
			center: s[0],
			child:  s[1],
		})
	}

	return orbits
}

func TestGetCountOfOrbits(t *testing.T) {
	data := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}
	orbits := prepareTestData(data)

	actual := GetCountOfOrbits(orbits)
	expected := 42

	if actual != expected {
		t.Errorf("Got %d want %d", actual, expected)
	}
}

func TestGetOrbitalTransfers(t *testing.T) {
	data := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"}
	orbits := prepareTestData(data)

	actual := GetOrbitalTransfers(orbits)
	expected := 4

	if actual != expected {
		t.Errorf("Got %d want %d", actual, expected)
	}
}
