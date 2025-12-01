package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day1Part1() {
	input, _ := os.ReadFile("inputs/1.txt")

	dial := 50
	password := 0

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}
		direction := line[0]
		distance, _ := strconv.Atoi(line[1:])

		switch direction {
		case 'R':
			dial = (dial + distance) % 100
		case 'L':
			dial = (dial - distance) % 100
		}

		if dial == 0 {
			password += 1
		}
	}

	fmt.Println("Part 1 Password:", password)
}

func mod(a, b int) int {
	m := a % b
	if m == 0 {
		return 0
	}
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}

func Day1Part2() {
	input, _ := os.ReadFile("inputs/1.txt")

	prev_dial := 50
	dial := 50
	password := 0

	for _, line := range strings.Split(string(input), "\n") {
		prev_dial = dial
		if dial < 0 {
			panic("dial is negative")
		}
		if dial >= 100 {
			panic("dial is greater than 100")
		}

		if line == "" {
			continue
		}

		direction := line[0]
		distance, _ := strconv.Atoi(line[1:])
		fmt.Println("Current dial:", dial, "Direction:", string(direction), "Distance:", distance)

		if distance >= 100 {
			password += distance / 100
			distance = distance % 100
		}
		
		switch direction {
		case 'R':
			dial = (dial + distance)
		case 'L':
			dial = (dial - distance)
		}

		if dial >= 100 || (dial <= 0 && prev_dial > 0) {
			password += 1
			dial = mod(dial, 100)
		} else if dial <= 0 {
			dial = mod(dial, 100)
		}
	}

	fmt.Println("Part 2 Password:", password)
}
