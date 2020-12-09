package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	operation string
	argument  int
	executed  bool
}

// read instructions into array of struct
func parseInstructions(scanner *bufio.Scanner) []Instruction {
	var instructions []Instruction
	for scanner.Scan() {
		var instruction Instruction
		s := scanner.Text()
		if s == "" {
			continue
		}
		parts := strings.Split(s, " ")
		instruction.operation = parts[0]
		instruction.argument, _ = strconv.Atoi(parts[1])
		instruction.executed = false
		instructions = append(instructions, instruction)
	}

	return instructions
}

func readInstructions(instructions []Instruction) int {
	var acc = 0
	var found = false

	// traverse the instructions
	var index = 0
	for !found {
		if index == len(instructions) -1 {
			fmt.Printf("reached the last line!\n")
			found = true
			break
		}

		var instruction = &instructions[index]
		fmt.Printf("[%d] visiting [%v]\n", index, instruction)

		if instruction.executed == true {
			fmt.Printf("second visit! for index [%d]: [%v]\n", index, instruction)
			found = true
			acc = -1 // remove this line for pt1
			break

		} else if instruction.operation == "acc" {
			instruction.executed = true
			acc += instruction.argument
			index++

		} else if instruction.operation == "jmp" {
			instruction.executed = true
			index += instruction.argument

		} else if instruction.operation == "nop" {
			instruction.executed = true
			index++
		}
		fmt.Printf("[%d] visited. [%v]\n", index, instruction)
	}
	return acc
}

// traverse the instructions, and carefully shift one jmp or nop at the time

func traverseAndShiftInstructions(instructions []Instruction) int {
	var swapIndex = 0
	var acc = -1
	experiment := make([]Instruction, len(instructions))

	for index := range instructions {
		copy(experiment, instructions)
		var instruction = &experiment[index]
		if instruction.operation == "jmp" || instruction.operation == "nop" {
			swapIndex = index
			switch instruction.operation {
			case "nop":
				instruction.operation = "jmp"

			case "jmp":
				instruction.operation = "nop"
			}
			fmt.Printf("\n\nswapping %d, [%v]\n", swapIndex, experiment)
			acc := readInstructions(experiment)
			if acc != -1 {
				fmt.Printf("made it all the way! swapping %d: acc %d\n", swapIndex, acc)
				break
			} else {
				fmt.Printf("no success swapping %d %v\n", swapIndex, instruction)
				experiment := instructions
				fmt.Printf("instr: %v\n", instructions)
				fmt.Printf("exper: %v\n", experiment)
			}
		}
	}
	return acc
}

func main() {
	fmt.Printf("Day 8\n")

	file, err := os.Open("tiny.txt")

	if err != nil {
		fmt.Printf("error opening file\n")
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var instructions = parseInstructions(scanner)
	var accValue = traverseAndShiftInstructions(instructions)
	fmt.Printf("PART 1, acc: %d\n", accValue)
}
