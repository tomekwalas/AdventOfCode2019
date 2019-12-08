package day_6

type Orbit struct {
	center string
	child  string
}

type OrbitNode struct {
	name     string
	children []*OrbitNode
}

func GetCountOfOrbits(orbits []Orbit) int {
	universe := createUniverse(orbits)

	return countOrbits(universe, 0)
}

func GetOrbitalTransfers(orbits []Orbit) int {
	universe := createUniverse(orbits)
	santa := universe.findPathTo("SAN", []string{})
	you := universe.findPathTo("YOU", []string{})

	santa, you = removeCommonPath(santa, you)

	return len(santa) + len(you) - 4

}

func (u OrbitNode) findPathTo(name string, path []string) []string {
	path = append(path, u.name)
	if u.name == name {
		return path
	}

	if len(u.children) == 0 {
		return []string{}
	}
	var childPath []string
	for _, child := range u.children {
		childPath = append(childPath, child.findPathTo(name, childPath)...)
	}
	if len(childPath) == 0 {
		path = childPath
	} else {
		path = append(path, childPath...)
	}

	return path
}
func removeCommonPath(src, dest []string) ([]string, []string) {
	i := 0
	for {
		if src[i] == dest[i] && src[i+1] == dest[i+1] {
			i++
		} else {
			break
		}

	}
	src = src[i:]
	dest = dest[i:]

	return src, dest
}

func createUniverse(data []Orbit) OrbitNode {
	var center Orbit
	for _, orbit := range data {
		if orbit.center == "COM" {
			center = orbit
			break
		}
	}

	universe := createOrbit(center.center, data)

	return *universe
}

func countOrbits(node OrbitNode, counter int) int {
	childOrbitsCount := 0
	if len(node.children) == 0 {
		return counter
	}
	for _, child := range node.children {
		childOrbitsCount += countOrbits(*child, counter+1)
	}

	return childOrbitsCount + counter
}

func createOrbit(parent string, data []Orbit) *OrbitNode {
	children := findChildren(parent, data)
	node := OrbitNode{
		name:     parent,
		children: []*OrbitNode{},
	}

	if len(children) == 0 {
		return &node
	}

	for _, child := range children {
		childNode := createOrbit(child.child, data)
		node.children = append(node.children, childNode)
	}

	return &node
}

func findChildren(parent string, data []Orbit) []Orbit {
	var children []Orbit
	for _, orbit := range data {
		if orbit.center == parent {
			children = append(children, orbit)
		}
	}

	return children
}
