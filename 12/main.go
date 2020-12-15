package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type NavigationInstruction struct {
	instruction string
	value       int
}

type coordinates struct {
	east      int
	north     int
	direction int
}

func part2(instructions []NavigationInstruction) int {
	waypoint := coordinates{east: 10, north: 1, direction: 0}
	position := coordinates{east: 0, north: 0, direction: 0}

	for _, ins := range instructions {
		switch ins.instruction {
		case string('N'):
			waypoint.north += ins.value
		case string('S'):
			waypoint.north -= ins.value
		case string('E'):
			waypoint.east += ins.value
		case string('W'):
			waypoint.east -= ins.value
		case string('F'):
			position.east += ins.value * waypoint.east   // - position.east)
			position.north += ins.value * waypoint.north //- position.north)
			// waypoint.east = position.east + waypoint.east - position.east
			// waypoint.north = position.north + waypoint.north - position.north
		case string('L'):
			// look at east and north
			movement := ins.value / 90
			switch movement {
			case 1:
				east := waypoint.east
				waypoint.east = -waypoint.north
				waypoint.north = east
			case 2:
				waypoint.north = -waypoint.north
				waypoint.east = -waypoint.east
			case 3:
				north := waypoint.north
				waypoint.north = -waypoint.east
				waypoint.east = north
			}
		case string('R'):
			movement := ins.value / 90
			switch movement {
			case 1:
				north := waypoint.north
				waypoint.north = -waypoint.east
				waypoint.east = north
			case 2:
				waypoint.north = -waypoint.north
				waypoint.east = -waypoint.east
			case 3:
				east := waypoint.east
				waypoint.east = -waypoint.north
				waypoint.north = east
			}
		}
		fmt.Printf("ins %v, waypoint %v, position %v\n", ins, waypoint, position)
	}

	if position.east < 0 {
		position.east = -position.east
	}
	if position.north < 0 {
		position.north = -position.north
	}
	fmt.Printf("position: %v\n", position)
	return (position.east + position.north)
}

func part1(instructions []NavigationInstruction) int {
	position := coordinates{east: 0, north: 0, direction: 0}

	for _, ins := range instructions {
		switch ins.instruction {
		case string('N'):
			position.north += ins.value
		case string('S'):
			position.north -= ins.value
		case string('E'):
			position.east += ins.value
		case string('W'):
			position.east -= ins.value
		case string('F'):
			if position.direction == 0 {
				position.east += ins.value
			} else if position.direction == 90 {
				position.north -= ins.value
			} else if position.direction == 180 {
				position.east -= ins.value
			} else if position.direction == 270 {
				position.north += ins.value
			}
		case string('L'):
			position.direction -= ins.value
			if position.direction < 0 {
				position.direction += 360
			}
		case string('R'):
			position.direction += ins.value
			if position.direction >= 360 {
				position.direction -= 360
			}
		}
	}

	if position.east < 0 {
		position.east = -position.east
	}
	if position.north < 0 {
		position.north = -position.north
	}
	fmt.Printf("position: %v\n", position)
	return (position.east + position.north)
}

func parseInput() []NavigationInstruction {
	file, err := os.Open("input.txt")
	var instructions []NavigationInstruction

	if err != nil {
		fmt.Printf("error opening file\n")
		return instructions
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var navigationInstruction NavigationInstruction
		s := scanner.Text()
		if s == "" {
			continue
		}
		navigationInstruction.instruction = s[:1]
		navigationInstruction.value, _ = strconv.Atoi(s[1:])
		instructions = append(instructions, navigationInstruction)
	}
	return instructions
}

func main() {

	fmt.Printf("Day 12\n")

	instructions := parseInput()

	// fmt.Printf("part 1: %v\n", part1(instructions))
	fmt.Printf("part 2: %v\n", part2(instructions))
}
