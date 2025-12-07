package main

import (
	"os"
	"strings"
)

func buildBeam(s string) int {
	lines := strings.Split(s, "\n")
	var splits int

	for i := 1; i < len(lines); i++ {
		chars := strings.Split(lines[i], "")

		for j, char := range chars {
			if lines[i-1][j] == '|' || lines[i-1][j] == 'S' {
				if char == "^" {
					lines[i] = replaceAtIndex(lines[i], '|', j-1)
					lines[i] = replaceAtIndex(lines[i], '|', j+1)
					splits++
				} else {
					lines[i] = replaceAtIndex(lines[i], '|', j)
				}
			}
		}
	}

	return splits
}

func countParticles(l []string) int {
	numLines := make([][]int, len(l))
	total := 0

	for i := range l {
		chars := strings.Split(l[i], "")
		numLine := make([]int, len(chars))

		for j, char := range chars {
			switch char {
			case "^":
				if i > 0 {
					numLine[j-1] += numLines[i-1][j]
					numLine[j+1] += numLines[i-1][j]
				}
			case ".":
				if i > 0 {
					numLine[j] += numLines[i-1][j]
				}
			case "S":
				numLine[j] = 1
			}
		}
		numLines[i] = numLine
	}

	for _, val := range numLines[len(numLines)-1] {
		total += val
	}

	return total
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func main() {
	content, _ := os.ReadFile("day7Input.txt")
	result := buildBeam(string(content))

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	particleResult := countParticles(lines)

	println(result)
	println(particleResult)
}
