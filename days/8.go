package days

import (
	"maps"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
	Z int
}

func distance(a Point, b Point) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	dz := a.Z - b.Z

	return math.Sqrt(math.Pow(float64(dx), 2.0) + math.Pow(float64(dy), 2.0) + math.Pow(float64(dz), 2.0))
}

type PairDistance struct {
	A        Point
	B        Point
	Distance float64
}

func Day8Part1() {
	content, _ := os.ReadFile("inputs/8.txt")

	points := []Point{}

	for _, line := range strings.Split(string(content), "\n") {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		points = append(points, Point{X: x, Y: y, Z: z})
	}

	pairDistances := []PairDistance{}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dist := distance(points[i], points[j])
			pairDistances = append(pairDistances, PairDistance{A: points[i], B: points[j], Distance: dist})
		}
	}

	sort.Slice(pairDistances, func(i, j int) bool {
		return pairDistances[i].Distance < pairDistances[j].Distance
	})

	circuitId := 0
	circuits := map[Point]int{}

	circuits[pairDistances[0].A] = circuitId
	circuits[pairDistances[0].B] = circuitId
	circuitId++

	// for i := 1; i < 10; i++ {
	for i := 1; i < 1000; i++ {
		A := pairDistances[i].A
		B := pairDistances[i].B

		circuitIdA, aExists := circuits[A]
		circuitIdB, bExists := circuits[B]

		if aExists && bExists {
			if circuitIdA != circuitIdB {
				for point, cid := range circuits {
					if cid == circuitIdB {
						circuits[point] = circuitIdA
					}
				}
			}
		} else if aExists {
			circuits[B] = circuitIdA
		} else if bExists {
			circuits[A] = circuitIdB
		} else {
			circuits[A] = circuitId
			circuits[B] = circuitId
			circuitId++
		}
	}

	cuircuitCounts := map[int]int{}

	for _, circuitId := range circuits {
		_, exists := cuircuitCounts[circuitId]
		if exists {
			cuircuitCounts[circuitId]++
		} else {
			cuircuitCounts[circuitId] = 1
		}
	}

	keys := make([]int, 0, len(cuircuitCounts))
	for k := range cuircuitCounts {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return cuircuitCounts[keys[i]] > cuircuitCounts[keys[j]]
	})

	result := cuircuitCounts[keys[0]]

	for i := 1; i < 3; i++ {
		result *= cuircuitCounts[keys[i]]
	}

	println(result)
}

func Day8Part2() {
	content, _ := os.ReadFile("inputs/8.txt")

	points := []Point{}

	for _, line := range strings.Split(string(content), "\n") {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		points = append(points, Point{X: x, Y: y, Z: z})
	}

	pairDistances := []PairDistance{}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dist := distance(points[i], points[j])
			pairDistances = append(pairDistances, PairDistance{A: points[i], B: points[j], Distance: dist})
		}
	}

	sort.Slice(pairDistances, func(i, j int) bool {
		return pairDistances[i].Distance < pairDistances[j].Distance
	})

	circuitId := 0
	circuits := map[Point]int{}

	circuits[pairDistances[0].A] = circuitId
	circuits[pairDistances[0].B] = circuitId
	circuitId++

	for i := 1; i < len(pairDistances); i++ {
		A := pairDistances[i].A
		B := pairDistances[i].B

		circuitIdA, aExists := circuits[A]
		circuitIdB, bExists := circuits[B]

		if aExists && bExists {
			if circuitIdA != circuitIdB {
				for point, cid := range circuits {
					if cid == circuitIdB {
						circuits[point] = circuitIdA
					}
				}
			}
		} else if aExists {
			circuits[B] = circuitIdA
		} else if bExists {
			circuits[A] = circuitIdB
		} else {
			circuits[A] = circuitId
			circuits[B] = circuitId
			circuitId++
		}

		distinctCircuits := []int{}
		for circuitId := range maps.Values(circuits) {
			if !slices.Contains(distinctCircuits, circuitId) {
				distinctCircuits = append(distinctCircuits, circuitId)
				if len(distinctCircuits) > 2 {
					break
				}
			}
		}
		if len(distinctCircuits) == 1 && len(circuits) == len(points) {
			println(A.X * B.X)
			break
		}
	}
}
