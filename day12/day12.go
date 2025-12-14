package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Present struct {
	id     int
	area   int
	layout []string
}

type Grid struct {
	id       int
	sizeX    int
	sizeY    int
	area     int
	presents []int
}

func getTotals(presents map[int]Present, presentCounts []int) (totalPresents, totalArea int) {
	for id, presentCount := range presentCounts {
		totalArea += presentCount * presents[id].area
		totalPresents += presentCount
	}
	return totalPresents, totalArea
}

func getInputs(presentLines [][]string, gridLines []string) (map[int]Present, []Grid) {
	presents := make(map[int]Present)

	for _, presentLine := range presentLines {
		present := Present{}
		id, _ := strconv.Atoi(strings.TrimRight(presentLine[0], ":"))
		present.id = id
		present.layout = presentLine[1:]
		present.area = 0

		for _, line := range present.layout {
			present.area += strings.Count(line, "#")
		}
		presents[present.id] = present
	}

	grids := []Grid{}
	for i, gridLine := range gridLines {
		line := Grid{id: i}
		dimensionsAndCounts := strings.Split(gridLine, ":")
		dimensions := dimensionsAndCounts[0]
		counts := dimensionsAndCounts[1]

		fmt.Sscanf(dimensions, "%dx%d", &line.sizeX, &line.sizeY)
		line.area = line.sizeX * line.sizeY

		presentCounts := strings.Fields(counts)
		var presentCount []int
		for _, s := range presentCounts {
			i, _ := strconv.Atoi(s)
			presentCount = append(presentCount, i)
		}

		line.presents = presentCount
		grids = append(grids, line)
	}

	return presents, grids
}

func splitLine(str string) []string {
	return strings.Split(strings.TrimSpace(str), "\n")
}

func getAllowedRegionTotal(presents map[int]Present, grids []Grid) int {
	count := 0
	for _, r := range grids {
		presentCount, presentArea := getTotals(presents, r.presents)
		if presentArea > r.area {
			continue
		}
		if presentCount*9 <= r.area {
			count++
			continue
		}
	}

	return count
}

func main() {
	content, _ := os.ReadFile("day12Input.txt")
	segments := strings.Split(string(content), "\n\n")

	presentPart := segments[:6]
	regionPart := segments[6]

	var segmentLines [][]string
	for _, part := range presentPart {
		splitSegment := splitLine(part)
		segmentLines = append(segmentLines, splitSegment)
	}

	regions := splitLine(regionPart)

	presents, grids := getInputs(segmentLines, regions)

	total := getAllowedRegionTotal(presents, grids)
	log.Printf("total allowed regions: %d", total)
}
