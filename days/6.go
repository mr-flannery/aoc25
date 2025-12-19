package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(a ...int) int {
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}

func product(a ...int) int {
	total := 1
	for _, v := range a {
		total *= v
	}
	return total
}

func Day6Part1() {
	content, _ := os.ReadFile("inputs/6.txt")

	numRows := [][]int{}
	opRow := []string{}

	for i, row := range strings.Split(string(content), "\n") {
		if i < 4 {
			numRow := []int{}
			for _, substring := range strings.Split(row, " ") {
				if substring == "" {
					continue
				}
				val, _ := strconv.Atoi(substring)
				numRow = append(numRow, val)
			}
			numRows = append(numRows, numRow)
		} else {
			for _, substring := range strings.Split(row, " ") {
				if substring == "" {
					continue
				}
				opRow = append(opRow, substring)
			}
		}
	}

	grandTotal := 0

	for i := 0; i < len(opRow); i++ {
		num1, num2, num3, num4 := numRows[0][i], numRows[1][i], numRows[2][i], numRows[3][i]
		switch opRow[i] {
		case "+":
			grandTotal += sum(num1, num2, num3, num4)
		case "*":
			grandTotal += product(num1, num2, num3, num4)
		default:
			panic("unknown operator")
		}
	}

	fmt.Println(grandTotal)
}

func Day6Part2() {
	content, _ := os.ReadFile("inputs/6.txt")

	rows := make([][]string, 5)

	for i, r := range strings.Split(string(content), "\n") {
		rows[i] = strings.Split(r, "")
	}

	for i := range rows {
		rows[i] = append(rows[i], " ")
	}

	total := 0
	currentOperand := ""
	nums := []int{}

	for i := 0; i < len(rows[0]); i++ {
		if rows[0][i] == " " && rows[1][i] == " " && rows[2][i] == " " && rows[3][i] == " " && rows[4][i] == " " {
			switch currentOperand {
			case "+":
				total += sum(nums...)
			case "*":
				total += product(nums...)
			default:
				panic("missing operator")
			}
			nums = []int{}
			currentOperand = ""
		} else {
			if currentOperand == "" {
				currentOperand = rows[4][i]
			}
			val, err := strconv.Atoi(strings.Trim(rows[0][i]+rows[1][i]+rows[2][i]+rows[3][i], " "))
			if err != nil {
				panic(err)
			}
			nums = append(nums, val)
		}
	}

	fmt.Println(total)
}
