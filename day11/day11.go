package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var testInput = `svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`

func parseGraph(input string) map[string][]string {
	graph := make(map[string][]string)
	lines := strings.SplitSeq(strings.TrimSpace(input), "\n")

	for line := range lines {
		parts := strings.Split(line, ":")
		from := strings.TrimSpace(parts[0])
		neighbors := strings.Fields(strings.TrimSpace(parts[1]))
		graph[from] = neighbors
	}

	return graph
}

func countPaths(graph map[string][]string, start, target string) int {
	visited := make(map[string]bool)
	return dfs(graph, start, target, visited)
}

func dfs(graph map[string][]string, current, target string, visited map[string]bool) int {
	if current == target {
		return 1
	}

	visited[current] = true
	defer func() { visited[current] = false }()

	total := 0
	for _, next := range graph[current] {
		if next != "out" && visited[next] {
			continue
		}
		total += dfs(graph, next, target, visited)
	}

	return total
}

func countPathsMemo(graph map[string][]string, start, target string, required []string) int {
	// Map required nodes to bit positions
	reqIndex := map[string]int{}
	for i, r := range required {
		reqIndex[r] = i
	}

	// Memoization: map[(node, mask)] = count
	memo := make(map[[2]any]int)

	visited := make(map[string]bool)

	return dfsWithReq(graph, start, target, reqIndex, 0, visited, memo)
}

func dfsWithReq(
	graph map[string][]string,
	node, target string,
	reqIndex map[string]int,
	mask int,
	visited map[string]bool,
	memo map[[2]any]int) int {
	// Detect visiting required node → update mask
	if idx, ok := reqIndex[node]; ok {
		mask |= (1 << idx)
	}

	// If reached target, check if all bits are 1
	if node == target {
		fullMask := (1 << len(reqIndex)) - 1
		if mask == fullMask {
			return 1
		}
		return 0
	}

	// Memo key
	key := [2]any{node, mask}
	if val, ok := memo[key]; ok {
		return val
	}

	visited[node] = true
	defer func() { visited[node] = false }()

	total := 0
	for _, nxt := range graph[node] {

		// Allow revisiting only the final "out", not other nodes
		if nxt != target && visited[nxt] {
			continue
		}

		total += dfsWithReq(graph, nxt, target, reqIndex, mask, visited, memo)
	}

	memo[key] = total
	return total
}

func main() {
	content, _ := os.ReadFile("day11Input.txt")
	currentTime := time.Now()
	graph := parseGraph(string(content))

	start := "you"
	target := "out"

	count := countPaths(graph, start, target)
	fmt.Println("Paths from", start, "to", target+":", count)

	println("part1 Time: ", int(time.Since(currentTime).Microseconds()))

	currentTime = time.Now()
	start = "svr"
	required := []string{"dac", "fft"}

	countPart2 := countPathsMemo(graph, start, target, required)
	fmt.Println("Valid paths:", countPart2)

	println("part2 Time: ", int(time.Since(currentTime).Microseconds()))
}
