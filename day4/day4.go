package main

import (
	"fmt"
	"os"
	"strings"
)

type Grid []string

func (g Grid) inBounds(z complex128) bool {
	y, x := int(real(z)), int(imag(z))
	return y >= 0 && y < len(g) && x >= 0 && x < len(g[y])
}

func (g Grid) charAt(z complex128) byte {
	if !g.inBounds(z) {
		return '.'
	}
	return g[int(real(z))][int(imag(z))]
}

func (g Grid) setChar(z complex128, c byte) {
	if !g.inBounds(z) {
		return
	}
	y, x := int(real(z)), int(imag(z))
	row := []byte(g[y])
	row[x] = c
	g[y] = string(row)
}

var neighbors = []complex128{
	-1 - 1i, -1 + 0i, -1 + 1i,
	0 - 1i, 0 + 1i,
	1 - 1i, 1 + 0i, 1 + 1i,
}

func countNeighbors(grid Grid, z complex128) int {
	count := 0
	for _, direction := range neighbors {
		c := grid.charAt(z + direction)
		if c == '@' || c == 'X' {
			count++
		}
	}
	return count
}

func removeWeakRolls(grid Grid) (removed int, changed bool) {
	for y := range grid {
		for x := 0; x < len(grid[y]); x++ {
			z := complex(float64(y), float64(x))

			if grid.charAt(z) == '@' {
				if countNeighbors(grid, z) < 4 {
					grid.setChar(z, 'X')
					removed++
					changed = true
				}
			}
		}
	}
	return removed, changed
}

func clearX(grid Grid) {
	for i := range grid {
		grid[i] = strings.ReplaceAll(grid[i], "X", ".")
	}
}

func recursiveRemoval(g Grid) int {
	total := 0
	for {
		removed, changed := removeWeakRolls(g)
		if !changed {
			return total
		}
		total += removed
		clearX(g)
	}
}

func main() {
	content, _ := os.ReadFile("day4Input.txt")
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")

	result := recursiveRemoval(lines)
	fmt.Println("Total rolls removed:", result)
}
