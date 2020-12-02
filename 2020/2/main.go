package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("sample.txt")

	if err != nil {
		fmt.Printf("error opening file")
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var input []int

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Printf("oops, error")
		}
		input = append(input, i)
	}
	fmt.Printf("Day 2, going steady \n")
	fmt.Printf("input: %v \n", input)
}
