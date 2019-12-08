package main

import "fmt"

var left = []string{"R8","U5","L5","D3"}
var right = []string{"U7","R6","D4","L4"}

var coordsLeft []Coords
var coordsRight []Coords

type Coords struct {
	x, y int
}


func parse(input []string, coordsLeft []Coords) {
	// parse and create array of coords
	for i := 0; i < 2; i++ {
		fmt.Println("first loop", i, len(input))
		var instruction = input[i]
		fmt.Println("instruction", instruction)
		var direction = string(instruction[0])
		var steps = int(instruction[1])
		fmt.Println("directions", direction, "steps", steps, "unparsed", instruction[1])
		for j := 0; j < steps; j++ {
			fmt.Println("here", i, "and", i)
			var x = 0
			var y = 0
			if (len(coordsLeft) > 0) {
				x = int(coordsLeft[len(coordsLeft)-1].x)
				y = int(coordsLeft[len(coordsLeft)-1].y)
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
			fmt.Println("new coord", newCoord)
			coordsLeft = append(coordsLeft, newCoord)
		}
	}
	fmt.Println("coords", coordsLeft)
}
func main() {
	fmt.Println("here we go")
	parse(left, coordsLeft)
	fmt.Println("coords: ", coordsLeft)
	// extend to two full arrays of cells (x,y)
	// find the common cells
	// measure how far they are from origo

}