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

func findCommonAffirmativeAnswer(scanner *bufio.Scanner) int {
	// each group is separated by blank line
	// each person, within the group is on a separate line
	var commonAffirmativeAnswers = ""
	var groupResponses []string

	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			var common = groupResponses[0]
			for _, response := range groupResponses {
				for _, char := range common {
					charString := string(char)
					if !strings.Contains(response, charString) {
						common = strings.ReplaceAll(common, charString, "")
					}
				}
			}
			commonAffirmativeAnswers += common

			groupResponses = make([]string, 0)
		} else {
			groupResponses = append(groupResponses, s)
		}
	}
	return len(commonAffirmativeAnswers)
}

func main() {
	fmt.Printf("Day 6\n")

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Printf("error opening file\n")
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	// var anyInGroupResponsesTotal = findAnyAffirmativeAnswer(scanner)
	// fmt.Printf("PART 1, number of right answers: %d\n", anyInGroupResponsesTotal)

	var commonInGroupResponsesTotal = findCommonAffirmativeAnswer(scanner)
	fmt.Printf("PART 2, number of right answers: %d\n", commonInGroupResponsesTotal)
}
