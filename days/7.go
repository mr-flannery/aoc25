package days

import (
	"os"
	"strconv"
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

func printGrid(grid [][]int) {
	for r := range grid {
		for c := range grid[r] {
			numString := strconv.Itoa(grid[r][c])
			if len(numString) < 2 {
				numString = " " + numString
			}
			print(numString)
		}
		println()
	}
}

func Day7Part2() {
	content, _ := os.ReadFile("inputs/7.txt")

	stringRows := [][]string{}
	for _, r := range strings.Split(string(content), "\n") {
		stringRows = append(stringRows, strings.Split(r, ""))
	}

	rows := make([][]int, len(stringRows))
	for i := range rows {
		rows[i] = make([]int, len(stringRows[i]))
	}

	for r := range stringRows {
		for c := range stringRows[r] {
			if stringRows[r][c] == "S" {
				rows[r][c] = -2
			} else if stringRows[r][c] == "^" {
				rows[r][c] = -1
			} else {
				rows[r][c] = 0
			}
		}
	}

	MAX_ROW_INDEX := len(rows) - 1
	MAX_COL_INDEX := len(rows[0]) - 1
	splits := 0

	for r := range rows {
		for c := range rows[r] {
			if rows[r][c] == -2 {
				rows[r+1][c] = 1
			}
			if r == MAX_ROW_INDEX && rows[r][c] > 0 {
				splits += rows[r][c]
			}
			if rows[r][c] > 0 {
				if r < MAX_ROW_INDEX {
					if rows[r+1][c] == -1 {
						if c > 0 {
							if rows[r+1][c-1] == -1 {
								panic("found a splitter in a place where a new beam should go")
							} else {
								rows[r+1][c-1] += rows[r][c]
							}
						}
						if c < MAX_COL_INDEX {
							if rows[r+1][c+1] == -1 {
								panic("found a splitter in a place where a new beam should go")
							} else {
								rows[r+1][c+1] += rows[r][c]
							}
						}
						// printGrid(rows)
					} else {
						rows[r+1][c] += rows[r][c]
						// printGrid(rows)
					}
				}
			}
		}
	}

	println(splits)
}
