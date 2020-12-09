package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
type Contains struct {
	number int
	color string
}

type Bag struct {
	color string
	contains []Contains
}

func parseRules(scanner *bufio.Scanner) []Bag {
	var bags []Bag
	for scanner.Scan() {
		var bag Bag
		s := scanner.Text()
		if s == "" {
			continue
		}
		parts := strings.Split(s, "contain ")
		bag.color = strings.Trim(parts[0], " bags")
		contains := strings.Split(parts[1], ", ")

		var contents []Contains
		// content: "1 bright white bag"
		// "no other bag"
		for _, content := range contains {
			var c Contains
			if content == "no other bags." {
				continue
			}
			splits := strings.Split(content, " ")
			c.number, _ = strconv.Atoi(splits[0])
			c.color = strings.Join(splits[1:len(splits)-1], " ")
			contents = append(contents, c)
		}
		bag.contains = contents
		bags = append(bags, bag)
	}

	return bags
}

func main() {
	fmt.Printf("Day 7\n")

	file, err := os.Open("tiny.txt")

	if err != nil {
		fmt.Printf("error opening file\n")
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var rules = parseRules(scanner)
	fmt.Printf("PART 1 %v\n", rules)
}
