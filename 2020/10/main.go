package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// traverse the list of joltage, starting from 0
// count the differences d
func traverseAdapters(adapters []int) int {
	var diff1 = 0
	var diff3 = 0

	var current = 0
	for _, adapter := range adapters {
		diff := adapter - current
		if diff == 1 {
			diff1++
		}
		if diff == 3 {
			diff3++
		}

		current = adapter
	}

	// Finally, your device's built-in adapter is always 3 higher than the highest adapter
	// so its rating is 22 jolts (always a difference of 3)
	diff3++
	return diff1 * (diff3)
}

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

	sort.Ints(numbers)
	return numbers
}

func main() {
	fmt.Printf("Day 10\n")

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Printf("error opening file\n")
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	numbers := parseInput(scanner)
	sum := traverseAdapters(numbers)
	fmt.Printf("PART 1, sum %d\n", sum)
}
