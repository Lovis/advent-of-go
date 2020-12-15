package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseInput() [][]string {
	file, err := os.Open("tiny.txt")
	var seatmap [][]string

	if err != nil {
		fmt.Printf("error opening file\n")
		return seatmap
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var row []string
		s := scanner.Text()
		if s == "" {
			continue
		}
		row = strings.Split(s, "")
		seatmap = append(seatmap, row)
	}
	return seatmap
}

func main() {

	fmt.Printf("Day 11\n")

	seatmap := parseInput()

	fmt.Printf("part 1: %v\n", seatmap)
}
