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

var answers = make(map[int]int)

func distinctWaysOfTraveling(adapters []int, index int) int {
	var count = 0
	if index == len(adapters)-1 {
		return 1
	}

	if answer, ok := answers[adapters[index]]; ok {
		return answer
	}

	for j, value := range adapters[index+1:] {
		if value-adapters[index] <= 3 {
			count += distinctWaysOfTraveling(adapters, j+index+1)
		}
	}
	answers[adapters[index]] = count

	return count
}

func dp(adapters []int) int {
	ans := map[int]int{0: 1}
	for _, a := range adapters {
		ans[a] = ans[a-1] + ans[a-2] + ans[a-3]
	}
	return ans[adapters[len(adapters)-1]]
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

	numbers := append([]int{0}, parseInput(scanner)...)
	// sum := traverseAdapters(numbers)
	// fmt.Printf("PART 1, sum %d\n", sum)

	variations := distinctWaysOfTraveling(numbers, 0)

	// Emil solution
	result := dp(numbers[1:])

	fmt.Printf("PART 2, variations %d\n", variations)
	fmt.Printf("PART 2, dp %d\n", result)

}
