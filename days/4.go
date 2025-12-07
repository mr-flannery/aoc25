package days

import (
	"fmt"
	"os"
	"strings"
)

func Day4Part1() {
	content, _ := os.ReadFile("inputs/4.txt")

	rows := strings.Split(string(content), "\n")

	movables := 0

	for r := 0; r < len(rows); r += 1 {
		for c := 0; c < len(rows[r]); c += 1 {
			if rows[r][c] == '@' {
				neighborScrolls := 0
				for dr := r - 1; dr <= r+1; dr += 1 {
					for dc := c - 1; dc <= c+1; dc += 1 {
						if (dr == r && dc == c) || dr < 0 || dc < 0 || dr >= len(rows) || dc >= len(rows[dr]) {
							continue
						}

						if rows[dr][dc] == '@' {
							neighborScrolls += 1
						}
					}
				}
				if neighborScrolls < 4 {
					movables += 1
				}
			}
		}
	}

	fmt.Println("Movable scrolls:", movables)
}

func Day4Part2() {
	content, _ := os.ReadFile("inputs/4.txt")
	rows := strings.Split(string(content), "\n")

	neighbors := make([][]int, len(rows))
	initialScrolls := 0

	for r := 0; r < len(rows); r += 1 {
		neighbors[r] = make([]int, len(rows[r]))
		for c := 0; c < len(rows[r]); c += 1 {
			neighbors[r][c] = 0
			if rows[r][c] == '@' {
				initialScrolls += 1
				for dr := r - 1; dr <= r+1; dr += 1 {
					for dc := c - 1; dc <= c+1; dc += 1 {
						if (dr == r && dc == c) || dr < 0 || dc < 0 || dr >= len(rows) || dc >= len(rows[dr]) {
							continue
						}

						if rows[dr][dc] == '@' {
							neighbors[r][c] += 1
						}
					}
				}
			}
		}
	}

	queue := [][]int{}

	// enqueue all scrolls with less than 4 neighbors
	for r := 0; r < len(rows); r += 1 {
		for c := 0; c < len(rows[r]); c += 1 {
			if rows[r][c] == '@' && neighbors[r][c] < 4 {
				queue = append(queue, []int{r, c})
			}
		}
	}

	remainingScrolls := initialScrolls
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		r, c := pos[0], pos[1]

		// decrement neighbor count on all adjacent scrolls by 1 and enqueue if they drop below 4
		for dr := r - 1; dr <= r+1; dr += 1 {
			for dc := c - 1; dc <= c+1; dc += 1 {
				if (dr == r && dc == c) || dr < 0 || dc < 0 || dr >= len(rows) || dc >= len(rows[dr]) {
					continue
				}

				if rows[dr][dc] == '@' {
					neighbors[dr][dc] -= 1
					if neighbors[dr][dc] == 3 {
						queue = append(queue, []int{dr, dc})
					}
				}
			}
		}
		// "remove" scroll
		remainingScrolls -= 1
		rows[r] = rows[r][:c] + "." + rows[r][c+1:] // maybe explicitly splitting this into a rune arr would've been cleaner 🤷
	}

	fmt.Println("Number of removed scrolls:", initialScrolls-remainingScrolls)
}
