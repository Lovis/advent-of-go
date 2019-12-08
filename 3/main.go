package main

import (
	"fmt"
	"strconv"
)

var left = []string{"R8","U5","L5","D3"}
var right = []string{"U7","R6","D4","L4"}

var coordsLeft []Coords
var coordsRight []Coords

type Coords struct {
	x, y int
}

func calcDistance(coords []Coords) (smallest int) {
	// traverse the list, find the smallest
	smallest = coords[0].x + coords[0].y
	for _, v := range coords {
		manhattan := v.x + v.y
		if (manhattan < smallest) {
			smallest = manhattan
		}
	}
	return smallest
}

func findIntersections(left []Coords, right []Coords) []Coords {
	var intersections []Coords
	for i, v := range left {
		fmt.Println(i, v, right[i])
		if (v.x == right[i].x && v.y == right[i].y) {
			intersections = append(intersections, v)
		}
	}
	return intersections
}

func parse(input []string, coords []Coords) ([]Coords) {
	// parse and create array of coords
	for i := 0; i < len(input); i++ {
		// fmt.Println("first loop", i, len(input))
		var instruction = input[i]
		// fmt.Println("instruction", instruction)
		var direction = string(instruction[0])
		steps, err := strconv.Atoi(string(instruction[1]))
		if err != nil {
			break
		}
		// fmt.Println("directions", direction, "steps", steps, "unparsed", instruction[1])
		for j := 0; j < steps; j++ {
			// fmt.Println("here", i, "and", i)
			var x = 0
			var y = 0
			if (len(coords) > 0) {
				x = int(coords[len(coords)-1].x)
				y = int(coords[len(coords)-1].y)
			}
			switch direction {
			case "R":
				x += 1
			case "U":
				y += 1
			case "L":
				x -= 1
			case "D":
				y -= 1
			}

			newCoord := Coords{x,y}
			// fmt.Println("new coord", newCoord)
			coords = append(coords, newCoord)
		}
	}
	return coords
}
func main() {
	fmt.Println("here we go")
	coordsLeft = parse(left, coordsLeft)
	fmt.Println("coordsLeft: ", coordsLeft)

	coordsRight = parse(right, coordsRight)
	fmt.Println("coordsRight: ", coordsRight)

	intersections := findIntersections(coordsLeft, coordsRight)
	fmt.Println("intersections", intersections)

	smallest := calcDistance(intersections)
	fmt.Println("smallest manhattan: ", smallest)
	// extend to two full arrays of cells (x,y)
	// find the common cells
	// measure how far they are from origo

}