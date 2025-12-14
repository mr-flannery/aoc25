package days

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day5Part1() {
	content, _ := os.ReadFile("inputs/5.txt")

	split := strings.Split(string(content), "\n\n")
	freshIngredientRangesString := split[0]
	ingredientIdsString := split[1]

	freshIngredientRanges := strings.Split(freshIngredientRangesString, "\n")
	ingredientIds := strings.Split(ingredientIdsString, "\n")
	
	freshIngIntIds := [][]int{}
	
	for _, line := range freshIngredientRanges {
		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		freshIngIntIds = append(freshIngIntIds, []int{start, end})
	}

	freshIngredientCount := 0
	for _, line := range ingredientIds {
		id, _ := strconv.Atoi(line)
		for _, rng := range freshIngIntIds {
			if id >= rng[0] && id <= rng[1] {
				freshIngredientCount++
				break
			}
		}
	}

	fmt.Println(freshIngredientCount)
}

func Day5Part2() {
	content, _ := os.ReadFile("inputs/5.txt")

	split := strings.Split(string(content), "\n\n")
	freshIngredientRangesString := split[0]

	freshIngredientRanges := strings.Split(freshIngredientRangesString, "\n")
	
	freshIngIntIds := [][]int{}
	
	for _, line := range freshIngredientRanges {
		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		updated := false
		for _, r := range freshIngIntIds {
			if (start >= r[0] && start <= r[1]) || (end >= r[0] && end <= r[1]) || (start <= r[0] && end >= r[1]) || (r[0] <= start && r[1] >= end) {
				r[0] = min(r[0], start)
				r[1] = max(r[1], end)
				updated = true
				break
			}
		}
		if !updated {
			freshIngIntIds = append(freshIngIntIds, []int{start, end})
		}
	}

	slices.SortFunc(freshIngIntIds, func(a []int, b []int) int {
		return a[0] - b[0]
	})

	for needsMerge := true; needsMerge;	{
		for i := 0; i < len(freshIngIntIds)-1; i++ {
			current := freshIngIntIds[i]
			next := freshIngIntIds[i+1] 
			
			if current[1] >= next[0] {
				freshIngIntIds[i][0] = min(current[0], next[0])
				freshIngIntIds[i][1] = max(current[1], next[1])
				freshIngIntIds[i+1] = []int{0, 0}
				i++
			}
		}
		
		lenBefore := len(freshIngIntIds)
		freshIngIntIds = slices.DeleteFunc(freshIngIntIds, func(a []int) bool {
			return a[0] == 0 && a[1] == 0
		})
		lenAfter := len(freshIngIntIds)

		needsMerge = lenBefore != lenAfter
	}

	totalFreshCount := 0
	for _, r := range freshIngIntIds {
		totalFreshCount += (r[1] - r[0] + 1)
	}

	fmt.Println(totalFreshCount)
}