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
		if index == len(instructions) {
			fmt.Printf("this is success!!\n")
			found = true
			break
		}

		var instruction = &instructions[index]
		fmt.Printf("visiting [%v]\n", instruction)

		if instruction.executed == true {
			fmt.Printf("second visit! for index [%d]: [%v]", index, instruction)
			found = true
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
		fmt.Printf("visited. [%v]\n", instruction)
	}
	return acc
}

func main() {
	fmt.Printf("Day 8\n")

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Printf("error opening file\n")
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var instructions = parseInstructions(scanner)
	var accValue = readInstructions(instructions)
	fmt.Printf("PART 1, acc: %d\n", accValue)
}
