package main

import (
	"os"
	"strconv"
	"strings"
)

var testInput = []string{
	"11-22", "95-115", "998-1012", "1188511880-1188511890", "222220-222224",
	"1698522-1698528", "446443-446449", "38593856-38593862", "565653-565659",
	"824824821-824824827", "2121212118-2121212124",
}

func getInvalidIds(i []string, matcher func(s string) bool) int {
	var sum int = 0
	for _, line := range i {
		if line != "" {
			stringSlice := strings.Split(line, "-")
			low, _ := strconv.ParseInt(stringSlice[0], 10, 64)
			high, _ := strconv.ParseInt(stringSlice[1], 10, 64)

			checkRange := make([]int, high-low+1)
			for i := range checkRange {
				num := i + int(low)
				str := strconv.Itoa(num)
				if matcher(str) {
					sum += num
				}
			}

		}
	}
	return sum
}

func isRepeated(s string) bool {
	half := s[0 : len(s)/2]

	return len(s)%2 == 0 && half+half == s
}

func isRepeated2(s string) bool {
	secondMatch := strings.Index((s + s)[1:], s) + 1

	return secondMatch < len(s)
}

func main() {
	content, _ := os.ReadFile("day2Input.txt")
	lines := strings.Split(strings.Trim(string(content), "\n"), ",")
	println(getInvalidIds(lines, isRepeated))
	println(getInvalidIds(lines, isRepeated2))
}
