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

type Result struct {
	Total int   // minimal total operations
	K     []int // counts per operation (length = m)
}

const IMPOSSIBLE_COUNT = 100_000

type Solver struct {
	parityMap map[int][]int // parity effect -> list of button masks
	buttons   []int         // operations as bitmasks
	memo      map[string]int
	m         int // number of indices
}

// NewSolver builds the parity map from ops
func NewSolver(ops [][]int, m int) *Solver {
	buttonMasks := make([]int, len(ops))
	for j, op := range ops {
		mask := 0
		for _, idx := range op {
			mask |= 1 << (m - 1 - idx) // reverse index
		}
		buttonMasks[j] = mask
	}

	parityMap := make(map[int][]int)
	total := 1 << len(buttonMasks)
	for mask := range total {
		parityEffect := 0
		for i := range buttonMasks {
			if (mask & (1 << i)) != 0 {
				parityEffect ^= buttonMasks[len(buttonMasks)-1-i]
			}
		}
		parityMap[parityEffect] = append(parityMap[parityEffect], mask)
	}

	return &Solver{
		parityMap: parityMap,
		buttons:   buttonMasks,
		memo:      make(map[string]int),
		m:         m,
	}
}

func listKey(v []int) string {
	parts := make([]string, len(v))
	for i, x := range v {
		parts[i] = fmt.Sprintf("%d", x)
	}
	return strings.Join(parts, ",")
}

func (s *Solver) Solve(target []int) int {
	key := listKey(target)
	if val, ok := s.memo[key]; ok {
		return val
	}

	allZero := true
	for _, x := range target {
		if x != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		s.memo[key] = 0
		return 0
	}

	for _, x := range target {
		if x < 0 {
			s.memo[key] = IMPOSSIBLE_COUNT
			return IMPOSSIBLE_COUNT
		}
	}

	// compute parity
	parity := 0
	for _, x := range target {
		parity = (parity << 1) | (x & 1)
	}

	opsMasks, ok := s.parityMap[parity]
	if !ok {
		s.memo[key] = IMPOSSIBLE_COUNT
		return IMPOSSIBLE_COUNT
	}

	minPresses := IMPOSSIBLE_COUNT

	for _, mask := range opsMasks {
		remaining := subtractButtonPress(target, mask, s.buttons, s.m)
		halved := halveAll(remaining)
		sub := s.Solve(halved)
		total := bitsCount(mask) + 2*sub
		if total < minPresses {
			minPresses = total
		}
	}

	s.memo[key] = minPresses
	return minPresses
}

func bitsCount(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return count
}

func halveAll(v []int) []int {
	out := make([]int, len(v))
	for i, x := range v {
		out[i] = x / 2
	}
	return out
}

func subtractButtonPress(target []int, mask int, buttons []int, m int) []int {
	out := append([]int(nil), target...)
	buttonIndex := len(buttons) - 1
	x := mask

	for x > 0 {
		if (x & 1) != 0 {
			val := buttons[buttonIndex]
			countIndex := m - 1
			for val > 0 {
				if val&1 != 0 {
					out[countIndex]--
				}
				val >>= 1
				countIndex--
			}
		}
		x >>= 1
		buttonIndex--
	}
	return out
}

func main() {
	contents, _ := os.ReadFile("day10Input.txt")
	result := processMachines(strings.TrimSpace(string(contents)))
	fmt.Println("Total minimum toggles:", result)

	// Part 2 - credit to veidom
	switchMachines := parseVectorInput(strings.TrimSpace(string(contents)))
	total := 0
	for i, sm := range switchMachines {
		println("Processing machine with index:", i)
		solver := NewSolver(sm.Toggles, len(sm.Target))
		minSteps := solver.Solve(sm.Target)
		total += minSteps
	}
	fmt.Println("Total min steps", total)

}
