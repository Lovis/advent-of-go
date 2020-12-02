package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PasswordPolicy struct {
	lower      int
	upper      int
	letter     string
	testPhrase string
}

func main() {

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Printf("error opening file")
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var input []PasswordPolicy

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		var policy PasswordPolicy

		policy.letter = strings.Trim(s[1], ":")
		policy.testPhrase = s[2]

		limits := strings.Split(s[0], "-")

		policy.lower, _ = strconv.Atoi(limits[0])
		policy.upper, _ = strconv.Atoi(limits[1])

		input = append(input, policy)

	}
	fmt.Printf("Day 2, going steady \n")
	fmt.Printf("input: %v \n", input)
}
