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
func readInstructions(scanner *bufio.Scanner) []Instruction {
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

func main() {
	fmt.Printf("Day 8\n")

	file, err := os.Open("tiny.txt")

	if err != nil {
		fmt.Printf("error opening file\n")
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var instructions = readInstructions(scanner)
	fmt.Printf("PART 1, instructions: %v\n", instructions)
}
