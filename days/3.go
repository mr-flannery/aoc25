package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day3Part1() {
	content, _ := os.ReadFile("inputs/3.txt")

	sum := 0

	for _, line := range strings.Split(string(content), "\n") {
		max := 0
		max_index := -1

		for i, char := range line {
			if val, _ := strconv.Atoi(string(char)); val > max {
				max = val
				max_index = i
			}
		}

		preMax := 0
		preMaxIndex := -1
		for i, char := range line[:max_index] {
			if val, _ := strconv.Atoi(string(char)); val > preMax {
				preMax = val
				preMaxIndex = i
			}
		}

		postMax := 0
		postMaxIndex := -1
		for i, char := range line[max_index+1:] {
			if val, _ := strconv.Atoi(string(char)); val > postMax {
				postMax = val
				postMaxIndex = i + max_index + 1
			}
		}

		candidate1 := 0
		if preMaxIndex != -1 {
			candidate1, _ = strconv.Atoi(string(line[preMaxIndex]) + string(line[max_index]))
		}

		candidate2 := 0
		if postMaxIndex != -1 {
			candidate2, _ = strconv.Atoi(string(line[max_index]) + string(line[postMaxIndex]))
		}

		if candidate1 > candidate2 {
			sum += candidate1
		} else {
			sum += candidate2
		}
	}

	fmt.Println("Part 1 Total Sum:", sum)
}

func Day3Part2() {
	content, _ := os.ReadFile("inputs/3.txt")

	sum := 0

	for _, line := range strings.Split(string(content), "\n") {
		indices := []int{}
		targetDigit := 9
		start := 0
		end := len(line) - 11
		for len(indices) < 12 {
			idx := strings.Index(line[start:end], strconv.Itoa(targetDigit))

			// when we're down to looking for a 1, we might as well take the last digits of the line and have a chance of finding something that is larger than 1
			if targetDigit == 1 {
				for i := len(line) - (12 - len(indices)); i < len(line); i++ {
					indices = append(indices, i)
				}
				break
			}

			if idx == -1 {
				targetDigit--
			} else {
				indices = append(indices, idx+start)
				start += idx + 1
				end += 1
				targetDigit = 9 // we have to reset the target digit, since we're extending our search window to indices we previously excluded
				// if the line has exactly the number of remaining digits as needed, just take them all
				if len(indices) < 12 && (len(line[start:])) == 12-len(indices) {
					for i := start; i < len(line); i++ {
						indices = append(indices, i)
					}
				}
			}
		}

		numStr := ""
		for _, idx := range indices {
			numStr += string(line[idx])
		}
 		val, _ := strconv.Atoi(numStr)
		fmt.Println("For line:", line, "constructed number:", numStr)
		sum += val
	}

	fmt.Println("Part 1 Total Sum:", sum)
}
