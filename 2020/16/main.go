package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseInput() {
	file, err := os.Open("tiny.txt")
	if err != nil {
		fmt.Printf("error opening file\n")
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s := scanner.Text()

		if s == "" {
			continue
		}

	}
}

func main() {
	fmt.Printf("hello, day 16\n")

	parseInput()
}
