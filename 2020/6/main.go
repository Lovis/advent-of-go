package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findAnyAffirmativeAnswer(scanner *bufio.Scanner) int {
	// each group is separated by blank line
	// each person, within the group is on a separate line
	var anyInGroupResponses = ""
	var groupResponses = ""

	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			// end of group
			anyInGroupResponses += groupResponses
			groupResponses = ""
		} else {
			for _, char := range s {
				charString := string(char)
				if !strings.Contains(groupResponses, charString) {
					groupResponses += charString
				}
			}
		}
	}
	return len(anyInGroupResponses)
}

func main() {
	fmt.Printf("Day 6\n")

	// var allInGroupResponses = ""

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Printf("error opening file\n")
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var anyInGroupResponsesTotal = findAnyAffirmativeAnswer(scanner)
	fmt.Printf("number of right answers: %d\n", anyInGroupResponsesTotal)
}
