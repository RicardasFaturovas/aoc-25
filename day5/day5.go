package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
)

func isInRange(value, min, max int) bool {
	return value >= min && value <= max
}

func getFreshAmount(ranges []string, ids []string) int {
	freshAmount := 0
	for _, id := range ids {
		for _, r := range ranges {
			var idNum, min, max int
			fmt.Sscanf(r, "%d-%d", &min, &max)
			fmt.Sscanf(id, "%d", &idNum)
			if isInRange(idNum, min, max) {
				freshAmount++
				break
			}
		}
	}

	return freshAmount
}

func getUniqueFreshAmount(ranges []string) int {
	freshAmount := 0
	uniqueRanges := [][]int{}

	slices.SortFunc(ranges, func(a, b string) int {
		var aMin, bMin int
		fmt.Sscanf(a, "%d-", &aMin)
		fmt.Sscanf(b, "%d-", &bMin)
		return cmp.Compare(aMin, bMin)
	})

	for i := 0; i <= len(ranges); i++ {
		if i >= len(ranges) {
			break
		}

		var r = ranges[i]
		var rMin, rMax int
		fmt.Sscanf(r, "%d-%d", &rMin, &rMax)

		if i+1 >= len(ranges) {
			break
		}

		newRange := []int{rMin, rMax}

		for _, nextR := range ranges[i+1:] {
			var nextMin, nextMax int
			fmt.Sscanf(nextR, "%d-%d", &nextMin, &nextMax)

			if nextMin <= rMax {
				rMax = max(rMax, nextMax)
				newRange = []int{rMin, rMax}
				// Skip the next range as it's merged
				i++
			}
		}
		uniqueRanges = append(uniqueRanges, newRange)
	}

	for _, uniqueR := range uniqueRanges {
		var (
			max = uniqueR[1]
			min = uniqueR[0]
		)
		freshAmount += (max - min + 1)
	}

	return freshAmount

}

func main() {
	content, _ := os.ReadFile("day5Input.txt")
	split := strings.Split(string(content), "\n\n")
	ranges := strings.Split(split[0], "\n")
	ids := strings.Split(split[1], "\n")

	result := getFreshAmount(ranges, ids)
	fmt.Println("Fresh amount:", result)

	uniqueResult := getUniqueFreshAmount(ranges)
	fmt.Println("Test Unique Fresh amount:", uniqueResult)

}
