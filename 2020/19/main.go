package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Rule: one to rule them all.
type Rule struct {
	index       int
	instruction string
	parsed      string
}

var rules = make(map[int]Rule)
var input []string

func getRule(rule int) []string {
	fmt.Printf("-- rule: %d\n", rule)
	var pattern []string = nil
	var curr = rules[rule]

	if len(curr.instruction) == 1 {
		curr.parsed = curr.instruction
		return append(pattern, curr.instruction)
	}
	subrules := strings.Split(curr.instruction, " | ")

	var subPattern []string
	for _, subrule := range subrules {
		var subRulePattern = ""

		ruleSet := strings.Split(subrule, " ")
		for _, sub := range ruleSet {
			next, _ := strconv.Atoi(sub)
			if rules[next].parsed != "" {
				subRulePattern += rules[next].parsed
			} else {
				nextRule := getRule(next)
				fmt.Printf("got fresh  %v\n", nextRule)
				parsed := rules[next]
				parsed.parsed = strings.Join(nextRule, "|")
				rules[next] = parsed
				subRulePattern += strings.Join(nextRule, " ")
			}
		}
		subPattern = append(subPattern, " "+subRulePattern)
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
			cursor = "input"
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
		fmt.Printf("[%d]: %v\n", key, value.instruction)
	}
	fmt.Printf("res2 %v\n", getRule(0))

	fmt.Printf("parsed rules: \n")
	for key, value := range rules {
		fmt.Printf("[%d]: ins: %s parsed: %v\n", key, value.instruction, value.parsed)
	}
}
