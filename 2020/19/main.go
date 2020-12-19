package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// try this
type Rule struct {
	index       int
	instruction string
	transcribed []string
}

var rules = make(map[int]Rule)
var input []string

func getRule(rule int) []string {
	fmt.Printf("-- rule: %d\n", rule)
	var pattern []string

	// handle sub
	subrules := strings.Split(rules[rule].instruction, " | ")
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
			if len(rules[next].instruction) == 1 {
				// assume a single character
				subRulePattern += rules[next].instruction
			} else {
				nextRule := getRule(next)
				fmt.Printf("subrule [%s], length: %d\n", nextRule, len(nextRule))
				subRulePattern += ("(" + strings.Join(nextRule, "|") + ")")
				fmt.Printf("done w sub rule pattern %v\n", subRulePattern)
			}
		}
		subPattern = append(subPattern, subRulePattern)
		// fmt.Printf("unjoined subpattern %v \n", subPattern)
	}
	fmt.Printf("joining subpattern %v, pattern %v\n", subPattern, pattern)
	pattern = append(pattern, strings.Join(subPattern, "|"))
	subPattern = nil

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
			var rule Rule
			rule.index = index
			rule.instruction = strings.Trim(parts[1], "\"")
			rules[index] = rule
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
		fmt.Printf("[%d]: %v\n", key, value)

	}
	fmt.Printf("getRule(5) %s\n", getRule(0))
}
