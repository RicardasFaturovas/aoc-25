package main

import (
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

var testInput = `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

type Box struct {
	X, Y, Z int
}

type Distance struct {
	Distance float64
	Box1     Box
	Box2     Box
}

func EuclideanDistance(b1, b2 Box) Distance {
	dx := b2.X - b1.X
	dy := b2.Y - b1.Y
	dz := b2.Z - b1.Z

	distanceSquared := (dx * dx) + (dy * dy) + (dz * dz)

	return Distance{
		Distance: math.Sqrt(float64(distanceSquared)),
		Box1:     b1,
		Box2:     b2,
	}
}

func getSortedDistances(points []Box) []Distance {
	distances := []Distance{}

	for i := range points {
		for j := i + 1; j < len(points); j++ {
			dist := EuclideanDistance(points[i], points[j])
			distances = append(distances, dist)
		}
	}

	distanceCompare := func(a, b Distance) int {
		return cmp.Compare(a.Distance, b.Distance)
	}

	slices.SortFunc(distances, distanceCompare)
	return distances
}

func buildCircuits(distances []Distance) [][]Box {
	circuits := [][]Box{}
	circuitIndexCache := map[Box]int{}

	for i, dist := range distances {
		box1Index, box1Exists := circuitIndexCache[dist.Box1]
		box2Index, box2Exists := circuitIndexCache[dist.Box2]

		if !box1Exists && !box2Exists {
			circuits = append(circuits, []Box{dist.Box1, dist.Box2})
			circuitIndexCache[dist.Box1] = len(circuits) - 1
			circuitIndexCache[dist.Box2] = len(circuits) - 1
		} else if box1Exists && !box2Exists {
			circuits[box1Index] = append(circuits[box1Index], dist.Box2)
			circuitIndexCache[dist.Box2] = box1Index
		} else if !box1Exists && box2Exists {
			circuits[box2Index] = append(circuits[box2Index], dist.Box1)
			circuitIndexCache[dist.Box1] = box2Index
		} else if box1Exists && box2Exists && box1Index != box2Index {
			circuits[box1Index] = append(circuits[box1Index], circuits[box2Index]...)

			for _, p := range circuits[box2Index] {
				circuitIndexCache[p] = box1Index
			}
			circuits[box2Index] = nil
		}

		// For part 2 stop, all circuits built
		if len(circuits[box1Index]) == 1000 {
			println(distances[i-1].Box1.X * distances[i-1].Box2.X)
			break
		}

	}
	nonNilCircuits := [][]Box{}
	for _, circuit := range circuits {
		if circuit != nil {
			nonNilCircuits = append(nonNilCircuits, circuit)
		}
	}
	circuits = nonNilCircuits

	return circuits
}

func sortCircuitsByLength(circuits [][]Box) [][]Box {
	slices.SortFunc(circuits, func(a, b []Box) int {
		return cmp.Compare(len(b), len(a))
	})
	return circuits
}

func main() {
	points := []Box{}
	content, _ := os.ReadFile("day8Input.txt")

	for line := range strings.FieldsSeq(string(content)) {
		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		points = append(points, Box{X: x, Y: y, Z: z})
	}

	sortedDistances := getSortedDistances(points)
	circuits := buildCircuits(sortedDistances[:1000])
	sortedCircuits := sortCircuitsByLength(circuits)

	result := len(sortedCircuits[0]) * len(sortedCircuits[1]) * len(sortedCircuits[2])
	fmt.Println("Result:", result)

	allCircuits := buildCircuits(sortedDistances)
	fmt.Println("Total Circuits:", len(allCircuits))
}
