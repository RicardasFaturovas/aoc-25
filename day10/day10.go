package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func minToggles(target int, toggles []int) int {
	type State struct {
		value int
		steps int
	}

	queue := []State{{0, 0}}
	visited := make(map[int]bool)
	visited[0] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.value == target {
			return current.steps
		}

		for _, toggle := range toggles {
			newState := current.value ^ toggle
			if !visited[newState] {
				visited[newState] = true
				queue = append(queue, State{newState, current.steps + 1})
			}
		}
	}

	return -1
}

type Machine struct {
	Target  int
	Toggles []int
}

type SwitchMachine struct {
	Target  []int
	Toggles [][]int
}

func lightsToBinary(input string) int {
	input = strings.Trim(input, "[]")
	result := 0

	for _, ch := range input {
		result <<= 1
		if ch == '#' {
			result |= 1
		}
	}
	return result
}

func positionsToBinary(input string, length int) int {
	input = strings.Trim(input, "()")
	if input == "" {
		return 0
	}

	result := 0
	positions := strings.SplitSeq(input, ",")

	for posStr := range positions {
		posStr = strings.TrimSpace(posStr)
		pos, err := strconv.Atoi(posStr)
		if err != nil {
			continue
		}

		if pos < 0 || pos >= length {
			continue
		}

		bitIndex := length - pos - 1
		result |= 1 << bitIndex
	}

	return result
}

func parseInput(input string) []Machine {
	var machines []Machine
	lines := strings.SplitSeq(input, "\n")
	for line := range lines {
		inputs := strings.Fields(line)
		target := lightsToBinary(inputs[0])

		toggles := inputs[1 : len(inputs)-1]
		var toggleBinaries []int
		for _, toggle := range toggles {
			toggleBinary := positionsToBinary(toggle, len(inputs[0])-2)
			toggleBinaries = append(toggleBinaries, toggleBinary)
		}

		machines = append(machines, Machine{Target: target, Toggles: toggleBinaries})
	}

	return machines
}

func processMachines(input string) int {
	machines := parseInput(input)
	total := 0
	for _, machine := range machines {
		minSteps := minToggles(machine.Target, machine.Toggles)
		total += minSteps
	}
	return total
}

func parseVectorInput(input string) []SwitchMachine {
	var machines []SwitchMachine

	lines := strings.SplitSeq(input, "\n")
	for line := range lines {
		inputs := strings.Fields(line)
		target := switchesToNumArr(inputs[len(inputs)-1])

		toggles := togglesToOps(inputs[1 : len(inputs)-1])

		machines = append(machines, SwitchMachine{Target: target, Toggles: toggles})
	}
	return machines
}

func switchesToNumArr(input string) []int {
	digits := strings.SplitSeq(strings.Trim(input, "{}"), ",")
	result := []int{}

	for ch := range digits {
		val, _ := strconv.Atoi(strings.TrimSpace(ch))
		result = append(result, val)
	}
	return result
}

func togglesToOps(toggles []string) [][]int {
	var ops [][]int
	for _, toggle := range toggles {
		toggle = strings.Trim(toggle, "()")
		positions := strings.SplitSeq(toggle, ",")
		var op []int
		for posStr := range positions {
			posStr = strings.TrimSpace(posStr)
			pos, _ := strconv.Atoi(posStr)
			op = append(op, pos)
		}
		ops = append(ops, op)
	}
	return ops
}

func main() {
	contents, _ := os.ReadFile("day10Input.txt")
	result := processMachines(strings.TrimSpace(string(contents)))

	fmt.Println("Total minimum toggles:", result)
	// Part 2 - skill issue
}
