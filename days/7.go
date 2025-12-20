package days

import (
	"os"
	"slices"
	"strings"
)

func Day7Part1() {
	content, _ := os.ReadFile("inputs/7.txt")

	rows := [][]string{}
	for _, r := range strings.Split(string(content), "\n") {
		rows = append(rows, strings.Split(r, ""))
	}

	MAX_ROW_INDEX := len(rows) - 1
	MAX_COL_INDEX := len(rows[0]) - 1
	splits := 0

	for r := range rows {
		for c := range rows[r] {
			if rows[r][c] == "S" {
				rows[r+1][c] = "|"
			}
			if rows[r][c] == "|" {
				if r < MAX_ROW_INDEX {
					if rows[r+1][c] == "^" {
						splits++
						if c > 0 {
							if rows[r+1][c-1] == "^" {
								panic("found a splitter in a place where a new beam should go")
							} else {
								rows[r+1][c-1] = "|"
							}
						}
						if c < MAX_COL_INDEX {
							if rows[r+1][c+1] == "^" {
								panic("found a splitter in a place where a new beam should go")
							} else {
								rows[r+1][c+1] = "|"
							}
						}
					} else {
						rows[r+1][c] = "|"
					}
				}
			}
		}
	}

	println(splits)
}

type Timeline struct {
	grid [][]string
	row  int
	col  int
}

func deepCopy(original [][]string) [][]string {
	if original == nil {
		return nil
	}

	copied := make([][]string, len(original))
	for i := range original {
		copied[i] = append([]string(nil), original[i]...)
	}
	return copied
}

func Day7Part2() {
	content, _ := os.ReadFile("inputs/7.txt")

	rows := [][]string{}
	for _, r := range strings.Split(string(content), "\n") {
		rows = append(rows, strings.Split(r, ""))
	}

	MAX_ROW_INDEX := len(rows) - 1
	MAX_COL_INDEX := len(rows[0]) - 1
	completedTimelines := 0

	c := slices.Index(rows[0], "S")
	rows[1][c] = "|"

	timelines := []Timeline{{
		grid: rows,
		row:  1,
		col:  c,
	}}

	// this is slow as fuck, implement the reddit thing instead
	// S becomes -2, ^ becomes -1, the initial beam gets a 1, everything becomes a 0
	for len(timelines) > 0 {
		current := timelines[0]
		timelines = timelines[1:]

		r := current.row
		c := current.col
		grid := current.grid

		if r == MAX_ROW_INDEX {
			completedTimelines++
		} else {
			if grid[r+1][c] == "^" {
				if c > 0 {
					cpy := deepCopy(grid)
					cpy[r+1][c-1] = "|"
					timelines = append(timelines, Timeline{
						grid: cpy,
						row:  r + 1,
						col:  c - 1,
					})
				}
				if c < MAX_COL_INDEX {
					cpy := deepCopy(grid)
					cpy[r+1][c+1] = "|"
					timelines = append(timelines, Timeline{
						grid: cpy,
						row:  r + 1,
						col:  c + 1,
					})
				}
			} else {
				grid[r+1][c] = "|"
				timelines = append(timelines, Timeline{
					grid: grid,
					row:  r + 1,
					col:  c,
				})
			}
		}
	}
	println(completedTimelines)
}
