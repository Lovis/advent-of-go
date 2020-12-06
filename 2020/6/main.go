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
	var groupResponses = ""
	var inProcess = false

	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			// end of group
			// fmt.Printf("summing, group unique: [%v]\n", groupResponses)
			commonAffirmativeAnswers += groupResponses
			groupResponses = ""
			inProcess = false
		} else if groupResponses == "" && !inProcess {
			groupResponses = s
			// fmt.Printf("\nfirst in new group: [%v]\n", groupResponses)
		} else {
			// groupResponses: abcx
			// s: aex
			// traverse groupResponses, if not existing
			// fmt.Printf("traversing: group: [%v], s [%s]\n", groupResponses, s)
			inProcess = true
			for _, char := range groupResponses {
				charString := string(char)
				if !strings.Contains(s, charString) {
					groupResponses = strings.ReplaceAll(groupResponses, charString, "")
				}
			}
		}
	}
	return len(commonAffirmativeAnswers)
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

	// var anyInGroupResponsesTotal = findAnyAffirmativeAnswer(scanner)
	// fmt.Printf("PART 1, number of right answers: %d\n", anyInGroupResponsesTotal)

	var commonInGroupResponsesTotal = findCommonAffirmativeAnswer(scanner)
	fmt.Printf("PART 2, number of right answers: %d\n", commonInGroupResponsesTotal)
}
