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

// traverse the list count the number of options:
//
func distinctWaysOfTraveling(adapters []int) int {

	// 1. convert to map for access by value
	var adapterMap = make(map[int]int)

	for _, a := range adapters {
		adapterMap[a] = a
	}

	var ways []int
	for value := range adapterMap {
		// for every index
		// 7 8 9 10
		var options []int
		if _, ok := adapterMap[value+1]; ok {
			options = append(options, value+1)
		}
		if _, ok := adapterMap[value+2]; ok {
			options = append(options, value+2)
		}
		if _, ok := adapterMap[value+3]; ok {
			options = append(options, value+3)
		}

		if value != adapters[len(adapters)-1] && len(options) > 1 {
			ways = append(ways, options...)
			fmt.Printf("round: %d, adding: %d total %d\n", value, options, ways)
		}

	}

	return len(ways) + 1
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

	file, err := os.Open("mini.txt")

	if err != nil {
		fmt.Printf("error opening file\n")
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	numbers := parseInput(scanner)
	// sum := traverseAdapters(numbers)
	// fmt.Printf("PART 1, sum %d\n", sum)

	variations := distinctWaysOfTraveling(numbers)
	fmt.Printf("PART 2, variations %d\n", variations)
}
