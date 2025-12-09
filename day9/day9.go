package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func area(p1, p2 Point) int {
	width := abs(p2.X-p1.X) + 1
	height := abs(p2.Y-p1.Y) + 1
	return width * height
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Rectangle struct {
	P1   Point
	P2   Point
	Area int
}

func getRectangles(points []Point) []Rectangle {
	rectangles := []Rectangle{}

	for i := range points {
		for j := i + 1; j < len(points); j++ {
			area := area(points[i], points[j])
			rectangles = append(rectangles, Rectangle{P1: points[i], P2: points[j], Area: area})
		}
	}

	return rectangles
}

func getSides(points []Point) [][2]Point {
	sides := [][2]Point{}
	for i, p := range points {
		nextP := points[0]
		if i+1 < len(points) {
			nextP = points[i+1]
		}
		sides = append(sides, [2]Point{p, nextP})
	}
	return sides
}

func inRange(a1, a2, b1, b2 int) bool {
	return !(a1 <= b1 && a1 <= b2 && a2 <= b1 && a2 <= b2) &&
		!(a1 >= b1 && a1 >= b2 && a2 >= b1 && a2 >= b2)
}

func sidesIntersect(p1 Point, p2 Point, sides [][2]Point) bool {
	result := false
	for _, side := range sides {
		if inRange(side[0].Y, side[1].Y, p1.Y, p2.Y) &&
			inRange(side[0].X, side[1].X, p1.X, p2.X) {
			result = true
			break
		}
	}

	return result
}

func main() {
	content, _ := os.ReadFile("day9Input.txt")
	lines := strings.Fields(string(content))
	points := []Point{}

	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, Point{X: x, Y: y})
	}

	rectangles := getRectangles(points)
	slices.SortFunc(rectangles, func(a, b Rectangle) int {
		return -cmp.Compare(a.Area, b.Area)
	})
	// part 1
	fmt.Println(rectangles[0].Area)

	sides := getSides(points)
	for _, rect := range rectangles {
		if !sidesIntersect(rect.P1, rect.P2, sides) {
			// part 2
			fmt.Println(rect.Area)
			break
		}

	}

}
