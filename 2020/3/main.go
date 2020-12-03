package main

import (
	"bufio"
	"fmt"
	"os"
)

// ..#...##
// #.#..###
// 2d matrix
// array of strings
// for _, row in input
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

	// path => asaaasssaass until length of array
	// [0][0], [1, 3], [2, 6]
	var count = 0
	for y := 0; y < len(pattern); y++ {
		if pattern[y][(3*y)%len(pattern[y])] == '#' {
			count++
		}
	}
	fmt.Printf("counted %d \n", count)
}
