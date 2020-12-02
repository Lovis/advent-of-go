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
	var count = 0
	// part 1
	// 1-3 a: abcde
	// for _, policy := range input {
	// 	// split on policy.letter
	// 	occurrencies := len(strings.Split(policy.testPhrase, policy.letter)) - 1
	// 	// fmt.Printf("policy %v, occurrencies %v\n", policy, occurrencies)
	// 	if occurrencies >= policy.lower && occurrencies <= policy.upper {
	// 		count++
	// 	}
	// }

	// part 2
	for _, policy := range input {
		// split on policy.letter 1-3 a abcde
		// => ac
		// strings.Split('ac', 'a')
		var first = (string(policy.testPhrase[policy.lower-1]) == policy.letter)
		var second = (string(policy.testPhrase[policy.upper-1]) == policy.letter)
		// fmt.Printf("policy %v, occurrencies %v\n", policy, occurrencies)
		if first == second {
			continue
		}
		if first || second {
			count++
		}
	}
	fmt.Printf("counted: %d\n", count)
}
