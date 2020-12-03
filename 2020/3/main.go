package main

import (
	"bufio"
	"fmt"
	"os"
)

func evenSteps(pattern []string) int {
	var count = 0
	for y := 0; y < len(pattern); y++ {
		if pattern[y][(3*y)%len(pattern[y])] == '#' {
			count++
		}
	}
	return count
}

func main() {
	fmt.Printf("Day 3 is here\n")

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Printf("error opening file")
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var pattern []string

	for scanner.Scan() {
		s := scanner.Text()
		pattern = append(pattern, s)

	}

	count := evenSteps(pattern)

	fmt.Printf("counted %d \n", count)
}
