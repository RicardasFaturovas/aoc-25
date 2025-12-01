package main

import (
	"os"
	"strconv"
	"strings"
)

var input = []string{
	"L68",
	"L30",
	"R48",
	"L5",
	"R60",
	"L55",
	"L1",
	"L99",
	"R14",
	"L82",
}

const (
	Min = 0
	Max = 100
)

func getPassword(i []string) int {
	var position int64 = 50
	var zeroCount int = 0

	for _, line := range i {
		if line != "" {
			var direction = string(line[0])
			var steps = line[1:]
			stepNumber, _ := strconv.ParseInt(steps, 10, 32)

			switch direction {
			case "L":
				position -= (stepNumber % Max)
			case "R":
				position += (stepNumber % Max)
			}
			position = (position + Max) % Max

			if position == 0 {
				zeroCount++
			}
		}
	}

	return zeroCount
}

func getPassword2(i []string) int {
	var position int64 = 50
	var zeroCount int = 0

	for _, line := range i {
		if line != "" {
			var direction = string(line[0])
			var steps = line[1:]
			var startingPosition = position

			stepNumber, _ := strconv.ParseInt(steps, 10, 32)
			quotient, remainder := divmod(stepNumber, Max)

			switch direction {
			case "L":
				position -= remainder
			case "R":
				position += remainder
			}

			zeroCount += int(quotient)
			if startingPosition != 0 && (position > Max || position < 0) {
				zeroCount++
			}

			position = (position + Max) % Max
			if position == 0 {
				zeroCount++
			}

		}
	}

	return zeroCount
}

func divmod(numerator, denominator int64) (quotient, remainder int64) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}

func main() {
	content, _ := os.ReadFile("day1Input.txt")
	lines := strings.Split(string(content), "\n")
	println(getPassword2(lines))
}
