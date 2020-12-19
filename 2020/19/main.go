package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rules = make(map[int]string)
var input []string

func testInput(s string) bool {
	for 
}

func parseInput() {
	file, err := os.Open("tiny.txt")
	if err != nil {
		fmt.Printf("error opening file\n")
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	cursor := "rules"

	for scanner.Scan() {
		s := scanner.Text()

		if s == "" {
			// switch to input
			cursor = "input"
			fmt.Printf("switching to %s\n", cursor)
			continue
		}

		switch cursor {
		case "rules":
			parts := strings.Split(s, ":")
			fmt.Printf("parts: %v\n", parts)
			index, _ := strconv.Atoi(parts[0])
			rules[index] = parts[1]
		case "input":
			input = append(input, s)
		}
	}
}

func main() {
	fmt.Printf("hello, day 19\n")

	parseInput()
	fmt.Printf("parsed input %v\n parsed rules: \n", input)
	for key, value := range rules {
		fmt.Printf("[%d]: %s\n", key, value)
	}
}
