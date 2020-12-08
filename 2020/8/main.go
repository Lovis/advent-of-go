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
		// fmt.Printf("index [%d] val [%v]\n", index, instructions[index])
		if instructions[index].executed == true {
			// fmt.Printf("second visit! for index [%d]: [%v]", index, instructions[index])
			found = true
			break
		} else if instructions[index].operation == "acc" {
			instructions[index].executed = true
			acc += instructions[index].argument
			index++
		} else if instructions[index].operation == "jmp" {
			instructions[index].executed = true
			index += instructions[index].argument
		} else if instructions[index].operation == "nop" {
			instructions[index].executed = true
			index++
		}

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
