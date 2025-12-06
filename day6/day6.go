package main

import (
	"os"
	"strconv"
	"strings"
)

type Grid [][]int

func getSum(input Grid, operations []string) int {
	var total int = 0

	for i, val := range input[0] {
		result := val
		for _, row := range input[1:] {
			switch operations[i] {
			case "*":
				result *= row[i]
			case "+":
				result += row[i]
			}
		}
		total += result
	}

	return total
}

func getInputs(s string) (Grid, []string) {
	var grid Grid = [][]int{}
	var operations []string
	lines := strings.Split(s, "\n")

	for i, line := range lines {
		if i == len(lines)-2 {
			operations = strings.Split(strings.ReplaceAll(line, " ", ""), "")
			break
		}

		splitLine := strings.Split(line, " ")

		var numLine []int
		for j := range splitLine {
			val := strings.ReplaceAll(splitLine[j], " ", "")
			if val != "" {
				num, _ := strconv.Atoi(strings.ReplaceAll(splitLine[j], " ", ""))
				numLine = append(numLine, num)
			}
		}
		grid = append(grid, numLine)
	}

	return grid, operations
}

func getResultPart2(s string) {
	lines := strings.Split(s, "\n")
	line := strings.Split(lines[0], "")
	operations := strings.Fields(lines[len(lines)-2])

	var verNumsArr [][]string
	var vertNums []string

	for i, l := range line {
		verticalNum := l
		for _, val := range lines[1 : len(lines)-2] {
			verticalNum += string(val[i])
		}

		if strings.ReplaceAll(verticalNum, " ", "") == "" {
			verNumsArr = append(verNumsArr, vertNums)
			vertNums = nil
		} else if i == len(line)-1 {
			vertNums = append(vertNums, verticalNum)
			verNumsArr = append(verNumsArr, vertNums)
			vertNums = nil
		} else {
			vertNums = append(vertNums, verticalNum)
		}
	}

	var total int
	for i, nArr := range verNumsArr {
		var result int
		switch operations[i] {
		case "*":
			result = 1
		case "+":
			result = 0
		}

		for _, val := range nArr {
			num, _ := strconv.Atoi(strings.ReplaceAll(val, " ", ""))
			switch operations[i] {
			case "*":
				result *= num
			case "+":
				result += num
			}
		}
		total += result
	}

	println(total)
}

func main() {
	content, _ := os.ReadFile("day6Input.txt")
	grid, operations := getInputs(string(content))
	result := getSum(grid, operations)

	println(result)
	getResultPart2(string(content))
}
