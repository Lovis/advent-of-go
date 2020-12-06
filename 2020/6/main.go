package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Day 6\n")

	var affirmativeAnswers = ""

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Printf("error opening file\n")
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	// each group is separated by blank line
	// each person, within the group is on a separate line

	var responses = ""
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			// end of group
			affirmativeAnswers += responses
			responses = ""
		} else {
			for _, char := range s {
				charString := string(char)
				if !strings.Contains(responses, charString) {
					responses += charString
				}
			}
		}
	}
	fmt.Printf("number of right answers: %d\n", len(affirmativeAnswers))

	// fmt.Printf("parsed answers: %v\n", affirmativeAnswers)
}
