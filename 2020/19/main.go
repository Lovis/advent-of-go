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

func getRule(rule int) string {
	fmt.Printf("get rule: %d len(%d)\n", rule, len(rules[rule]))
	var subs []string
	if len(rules[rule]) == 1 {
		// assume a single character
		fmt.Printf("got char; returning\n")
		return rules[rule]
	}
	// handle sub
	subrules := strings.Split(rules[rule], " | ")
	// 0: 1 2
	// 1: "a"
	// 2: 1 3 | 3 1
	// 3: "b"
	// 0 => a + ab|ba => aab | aba
	for i, subrule := range subrules {
		subs = append(subs, "")
		ruleSet := strings.Split(subrule, " ")
		for _, subsub := range ruleSet {
			fmt.Printf("subrule %d: subsub [%s]\n", i, subsub)
			next, _ := strconv.Atoi(subsub)
			subs[i] += getRule(next)
		}
	}
	return strings.Join(subs, "|")
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
			fmt.Printf("parts: %v\n", parts)
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
		fmt.Printf("[%d]: %s, len(%d)\n", key, value, len(value))

	}

	// fmt.Printf("len(rules[5]): %d\n", len(rules[5]))
	fmt.Printf("getRule(5) %s\n", getRule(3))
}
