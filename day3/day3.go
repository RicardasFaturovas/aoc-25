package main

import (
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

var testInput = []string{
	"987654321111111",
	"811111111111119",
	"234234234234278",
	"818181911112111",
}

func getMaxJoltage(joltages []string) int {
	var total int = 0

	for _, line := range joltages {
		if line != "" {
			var nums []int
			var maxJoltage int = 0

			for _, r := range line {
				val, _ := strconv.Atoi(string(r))
				nums = append(nums, val)
			}

			max := slices.Max(nums)
			maxIndex := slices.Index(nums, max)

			if maxIndex == len(nums)-1 {
				max = slices.Max(nums[:maxIndex])
				maxIndex = slices.Index(nums, max)
			}
			leftOver := nums[maxIndex+1:]
			secondMax := slices.Max(leftOver)

			maxJoltage = max*10 + secondMax

			total += maxJoltage
		}
	}
	return total
}

func getMaxJoltage2(joltages []string) int {
	var total int = 0

	for _, line := range joltages {
		if line != "" {
			var nums []int

			for _, r := range line {
				val, _ := strconv.Atoi(string(r))
				nums = append(nums, val)
			}

			maxJoltage := buildHighestJoltage(nums)
			total += maxJoltage
		}
	}
	return total
}

func buildHighestJoltage(nums []int) int {
	var total int = 0
	var maxIndex = 0

	for i := 12; i > 0; i-- {
		max := slices.Max(nums[:len(nums)-i+1])
		maxIndex = slices.Index(nums, max) + 1
		nums = nums[maxIndex:]

		total = total + max*int(math.Pow10(i-1))
	}

	return total
}

func main() {
	content, _ := os.ReadFile("day3Input.txt")
	lines := strings.Split(string(content), "\n")

	println(getMaxJoltage(lines))
	println(getMaxJoltage2(lines))
}
