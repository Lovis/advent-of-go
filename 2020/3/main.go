package main

import (
	"bufio"
	"fmt"
	"os"
)

func evenSteps(pattern []string, k int, x int) int {
	var count = 0
	for y := 0; y < len(pattern); y += x {
		if pattern[y][(k*y/x)%len(pattern[y])] == '#' {
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

	c1 := evenSteps(pattern, 1, 1)
	c3 := evenSteps(pattern, 3, 1)
	c5 := evenSteps(pattern, 5, 1)
	c7 := evenSteps(pattern, 7, 1)
	c2 := evenSteps(pattern, 1, 2)

	// fmt.Printf("counted %d %d %d %d %d\n", c1, c3, c5, c7, c2)
	fmt.Printf("counted %d \n", c1*c3*c5*c7*c2)
}
