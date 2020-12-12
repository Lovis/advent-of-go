package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func findContiguousSet(preamble []int, current int) int {
	// traverse the preamble list until the sum is equal to current
	var found = false
	var sum []int
	var currentSum int
	var smallest, largest int
	for i := 0; !found && i < len(preamble); i++ {

		first := preamble[i]
		currentSum = first
		sum = []int{first}
		for c := i + 1; currentSum < current && c < len(preamble); c++ {
			sum = append(sum, preamble[c])
			currentSum += preamble[c]

			if currentSum == current {
				fmt.Printf("found a sum matching: %v\n", sum)
				// now find smallest and largest
				smallest, largest = sum[0], sum[0]
				for _, cont := range sum {
					if smallest > cont {
						smallest = cont
					}
					if largest < cont {
						largest = cont
					}
				}
				found = true
				break
			}
		}
	}
	return smallest + largest
}

func verifyPreamble(preamble []int, current int) bool {
	// find the two numbers amongst the 5 previous elements
	// pluck the first element, find the missing value
	// fmt.Printf("preamble %v, number %d\n", preamble, current)
	var found = false
	for i := 0; !found && i < len(preamble); i++ {
		first := preamble[i]
		desiredPreamble := current - first
		for _, second := range preamble {
			if first == second {
				continue
			}

			if second == desiredPreamble {
				// fmt.Printf("found match: %d + %d = %d\n", first, second, current)
				found = true
				break
			}
		}
	}
	return found
}

func parsePreamble(scanner *bufio.Scanner, preambleLimit int) int {
	var numbers []int
	var mismatch int
	for scanner.Scan() {
		var current int
		s := scanner.Text()
		if s == "" {
			continue
		}
		current, _ = strconv.Atoi(s)
		numbers = append(numbers, current)
		if len(numbers) <= preambleLimit {
			continue
		}

		verified := verifyPreamble(numbers[len(numbers)-(preambleLimit+1):len(numbers)-1], current)

		if !verified {
			mismatch = current
			key := findContiguousSet(numbers, current)
			fmt.Printf("PART 2: found a contiguous set w sum: %d, mismatch: %d\n", key, mismatch)
			return key
		}
	}

	return -1
}

func main() {
	fmt.Printf("Day 9\n")

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Printf("error opening file\n")
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	fmt.Printf("PART 1+2, preamble: %d\n", parsePreamble(scanner, 25))
}
