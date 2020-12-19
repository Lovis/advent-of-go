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

func getRule(rule int) []string {
	fmt.Printf("-- rule: %d\n", rule)
	var pattern []string

	// handle sub
	subrules := strings.Split(rules[rule], " | ")
	// 0: 1 2
	// 1: "a"
	// 2: 1 3 | 3 1
	// 3: "b"
	// 0 => a + ab|ba => [a, ab|ba]
	var subPattern []string
	for _, subrule := range subrules {
		var subRulePattern = ""

		// " 1 3 " => [1,3]
		ruleSet := strings.Split(subrule, " ")
		for _, sub := range ruleSet {
			next, _ := strconv.Atoi(sub)
			if len(rules[next]) == 1 {
				// assume a single character
				subRulePattern += rules[next]
			} else {
				nextRule := getRule(next)
				fmt.Printf("subrule [%s], length: %d\n", nextRule, len(nextRule))
				subRulePattern += strings.Join(nextRule, "|")
			}
		}
		subPattern = append(subPattern, subRulePattern)
		fmt.Printf("unjoined subpattern %v \n", subPattern)
		pattern = append(pattern, strings.Join(subPattern, "|"))
		subPattern = nil
	}

	fmt.Printf("pattern len(%d) %v\n", len(pattern), pattern)
	return pattern
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
			parts := strings.Split(s, ": ")
			index, _ := strconv.Atoi(parts[0])
			rules[index] = strings.Trim(parts[1], "\"")
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
	fmt.Printf("getRule(5) %s\n", getRule(0))
}
