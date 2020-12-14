package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseInput(scanner *bufio.Scanner) []int {
	var numbers []int
	for scanner.Scan() {
		var current int
		s := scanner.Text()
		if s == "" {
			continue
		}
		current, _ = strconv.Atoi(s)
		numbers = append(numbers, current)
	}

	return numbers
}

func main() {
	fmt.Printf("Day 10\n")

	file, err := os.Open("tiny.txt")

	if err != nil {
		fmt.Printf("error opening file\n")
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	fmt.Printf("PART 1, joltage: %d\n", parseInput(scanner))
}
