package days

import (
	"math"
	"os"
	"strconv"
	"strings"
)

type Point2D struct {
	X int
	Y int
}

func Day9Part1() {
	content, _ := os.ReadFile("inputs/9.txt")

	lines := strings.Split(string(content), "\n")

	points := []Point2D{}

	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		points = append(points, Point2D{X: x, Y: y})
	}

	areaMax := 0

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			area := math.Abs(float64(points[i].X-points[j].X)+1) * math.Abs(float64(points[i].Y-points[j].Y)+1)
			if area > float64(areaMax) {
				areaMax = int(area)
			}
		}
	}

	println(areaMax)
}

func verticalLineIntersects(line [2]Point2D, pointA Point2D, pointB Point2D) bool {
	if line[0].X == line[1].X {
		return false // two vertical lines cannot interesect
	} else {
		for line_x := min(line[0].X, line[1].X); line_x <= max(line[0].X, line[1].X); line_x++ {
			if line_x > min(pointA.Y, pointB.Y) && line_x < max(pointA.Y, pointB.Y) {
				return true
			}
		}
		return false
	}
}

func horizontalLineIntersects(line [2]Point2D, pointA Point2D, pointB Point2D) bool {
	if line[0].Y == line[1].Y {
		return false // two horizontal lines cannot intersect
	} else {
		for line_y := min(line[0].Y, line[1].Y); line_y <= max(line[0].Y, line[1].Y); line_y++ {
			if line_y > min(pointA.X, pointB.X) && line_y < max(pointA.X, pointB.X) {
				return true
			}
		}
		return false
	}
}

func intersects(line [2]Point2D, pointA Point2D, pointB Point2D) bool {
	grid := make([][]string, max(pointA.Y, pointB.Y, line[0].Y, line[1].Y)+1)
	for i := range grid {
		grid[i] = make([]string, max(pointA.X, pointB.X, line[0].X, line[1].X)+1)
	}

	// draw square
	for r := min(pointA.Y, pointB.Y); r <= max(pointA.Y, pointB.Y); r++ {
		for c := min(pointA.X, pointB.X); c <= max(pointA.X, pointB.X); c++ {
			grid[r][c] = "S"
		}
	}

	// draw line
	if line[0].X == line[1].X {
		for r := min(line[0].Y, line[1].Y); r <= max(line[0].Y, line[1].Y); r++ {
			grid[r][line[0].X] = "L"
		}
	} else {
		for c := min(line[0].X, line[1].X); c <= max(line[0].X, line[1].X); c++ {
			grid[line[0].Y][c] = "L"
		}
	}

	// fill in dots
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == "" {
				grid[r][c] = "."
			}
		}
	}

	// v2: mathematical version
	// if line[0].X == line[1].X {
	// 	if line[0].X > min(pointA.X, pointB.X) && line[0].X < max(pointA.X, pointB.X) && ((line[0].Y > min(pointA.Y, pointB.Y) && line[0].Y < max(pointA.Y, pointB.Y)) || (line[1].Y > min(pointA.Y, pointB.Y) && line[1].Y < max(pointA.Y, pointB.Y)) || (min(line[0].Y, line[1].Y) <= min(pointA.Y, pointB.Y) && max(line[0].Y, line[1].Y) >= max(pointA.Y, pointB.Y))) {
	// 		return true
	// 	}
	// } else {
	// 	if line[0].Y > min(pointA.Y, pointB.Y) && line[0].Y < max(pointA.Y, pointB.Y) && ((line[0].X > min(pointA.X, pointB.X) && line[0].X < max(pointA.X, pointB.X)) || (line[1].X > min(pointA.X, pointB.X) && line[1].X < max(pointA.X, pointB.X)) || (min(line[0].X, line[1].X) <= min(pointA.X, pointB.X) && max(line[0].X, line[1].X) >= max(pointA.X, pointB.X))) {
	// 		return true
	// 	}
	// }
	// return false

	// v3:
	// if square is a line, check if the line has a point that is neither of the square's corners
	if pointA.X == pointB.X {
		return horizontalLineIntersects(line, pointA, pointB)
	} else if pointA.Y == pointB.Y {
		return verticalLineIntersects(line, pointA, pointB)
	} else {
		// if square is two-wide, treat it as two lines
		if math.Abs(float64(pointA.X)-float64(pointB.X)) == 1.0 {
			return horizontalLineIntersects(line, pointA, Point2D{X: pointA.X, Y: pointB.Y}) && horizontalLineIntersects(line, Point2D{X: pointB.X, Y: pointA.Y}, pointB)
		} else if math.Abs(float64(pointA.Y)-float64(pointB.Y)) == 1.0 {
			return verticalLineIntersects(line, pointA, Point2D{X: pointB.X, Y: pointA.Y}) && verticalLineIntersects(line, Point2D{X: pointA.X, Y: pointB.Y}, pointB)
		} else {
			// sqaure is three wide in at least one dimension
			// the line must have at least one point whose coords are in between the both corners x and y coords
			if line[0].X == line[1].X {
				for line_y := min(line[0].Y, line[1].Y); line_y <= max(line[0].Y, line[1].Y); line_y++ {
					if min(line[0].X, line[1].X) > min(pointA.X, pointB.X) && max(line[0].X, line[1].X) < max(pointA.X, pointB.X) && line_y > min(pointA.Y, pointB.Y) && line_y < max(pointA.Y, pointB.Y) {
						return true
					}
				}
				return false
			} else {
				for line_x := min(line[0].X, line[1].X); line_x <= max(line[0].X, line[1].X); line_x++ {
					if min(line[0].Y, line[1].Y) > min(pointA.Y, pointB.Y) && max(line[0].Y, line[1].Y) < max(pointA.Y, pointB.Y) && line_x > min(pointA.X, pointB.X) && line_x < max(pointA.X, pointB.X) {
						return true
					}
				}
				return false
			}
		}
	}

}

func Day9Part2() {
	content, _ := os.ReadFile("inputs/9.txt")

	lines := strings.Split(string(content), "\n")
	points := []Point2D{}

	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		points = append(points, Point2D{X: x, Y: y})
	}

	lineSegments := [][2]Point2D{}

	for i := 0; i < len(points); i += 1 {
		j := (i + 1)
		if i == len(points)-1 {
			j = 0
		}
		lineSegments = append(lineSegments, [2]Point2D{points[i], points[j]})
	}

	areaMax := 0

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			isValid := true

			for _, line := range lineSegments {
				if intersects(line, points[i], points[j]) {
					isValid = false
					break
				}
			}

			if !isValid {
				continue
			}

			area := math.Abs(float64(points[i].X-points[j].X)+1) * math.Abs(float64(points[i].Y-points[j].Y)+1)
			if area > float64(areaMax) {
				areaMax = int(area)
			}
		}
	}

	println(areaMax)
}
