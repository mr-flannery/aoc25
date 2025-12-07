package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2Part1() {
	content, _ := os.ReadFile("inputs/2.txt")

	ranges := strings.Split(string(content), ",")

	count := 0

	for _, r := range ranges {
		startAndEnd := strings.Split(r, "-")
		start, _ := strconv.Atoi(startAndEnd[0])
		end, _ := strconv.Atoi(startAndEnd[1])

		fmt.Println(start, end)
		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)
			if s[:len(s)/2] == s[len(s)/2:] {
				fmt.Println("Found matching number:", s)
				count += i
			}
		}
	}

	fmt.Println("Part 1 Total Count:", count)
}

func Day2Part2() {
	content, _ := os.ReadFile("inputs/2.txt")

	ranges := strings.Split(string(content), ",")

	count := 0

	for _, r := range ranges {
		startAndEnd := strings.Split(r, "-")
		start, _ := strconv.Atoi(startAndEnd[0])
		end, _ := strconv.Atoi(startAndEnd[1])

		fmt.Println(start, end)
		for numTocheck := start; numTocheck <= end; numTocheck++ {
			s := strconv.Itoa(numTocheck)
			for seqSize := 1; seqSize <= len(s)/2; seqSize++ {
				if len(s)%seqSize != 0 {
					continue
				}
				valid := true
				
				for i := seqSize; i <= len(s) - seqSize; i += seqSize {
					if valid && s[i-seqSize:i] != s[i:i+seqSize] {
						valid = false
						break
					}
				}

				if valid {
					fmt.Println("Found matching number:", s)
					count += numTocheck
					break
				}
			}
		}
	}

	fmt.Println("Part 2 Total Count:", count)
}	