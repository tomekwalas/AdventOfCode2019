package day_3

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	Left  = "L"
	Right = "R"
	Up    = "U"
	Down  = "D"
)

type Point struct {
	X, Y int
}

func (ps Point) GetDistance(pd Point) int {
	dx := math.Abs(float64(ps.X - pd.X))
	dy := math.Abs(float64(ps.Y - pd.Y))

	return int(dx + dy)
}

type Wire struct {
	value int
	name  string
}

func generateGrid(wires [][]string) map[string]Wire {
	grid := make(map[string]Wire)
	for num, wire := range wires {
		x := 0
		y := 0
		for _, part := range wire {
			direction := part[:1]
			number, _ := strconv.Atoi(part[1:])
			wireName := fmt.Sprintf("wire%d", num)
			for i := 0; i < number; i++ {
				switch direction {
				case Right:
					x++
					break
				case Up:
					y++
					break
				case Left:
					x--
					break
				case Down:
					y--
					break
				}
				coords := fmt.Sprintf("%d_%d", x, y)

				if grid[coords].name == "" {
					grid[coords] = Wire{value: 1, name: wireName}
				} else if grid[coords].name != wireName {
					grid[coords] = Wire{value: 2, name: wireName}
				}
			}
		}
	}
	return grid
}

func GetClosestIntersectionDistance(wires [][]string) int {
	grid := generateGrid(wires)

	base := Point{X: 0, Y: 0}
	min := 0
	for k, v := range grid {
		if v.value > 1 {
			data := strings.Split(k, "_")
			x, _ := strconv.Atoi(data[0])
			y, _ := strconv.Atoi(data[1])
			p := Point{X: x, Y: y}

			distance := base.GetDistance(p)
			if min == 0 {
				min = distance
			}
			if distance < min {
				min = distance
			}
		}
	}

	return min
}

func GetFewestStepsToIntersection(wires [][]string) int {
	grid := generateGrid(wires)
	var allSteps []int
	for k, v := range grid {
		if v.value > 1 {
			data := strings.Split(k, "_")
			x, _ := strconv.Atoi(data[0])
			y, _ := strconv.Atoi(data[1])
			p := Point{X: x, Y: y}
			steps := 0
			for _, wire := range wires {
				x := 0
				y := 0
				for _, part := range wire {
					if p.X == x && p.Y == y {
						break
					}
					direction := part[:1]
					number, _ := strconv.Atoi(part[1:])

					for i := 0; i < number; i++ {
						switch direction {
						case Right:
							x++
							break
						case Up:
							y++
							break
						case Left:
							x--
							break
						case Down:
							y--
							break
						}
						steps += 1
						if p.X == x && p.Y == y {
							break
						}
					}
				}
			}
			allSteps = append(allSteps, steps)
		}
	}
	min := allSteps[0]
	for _, step := range allSteps {
		if step < min {
			min = step
		}
	}
	return min
}
