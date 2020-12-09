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

func findShinyBags(bags []Bag) int {
	// use the []Bag
	// traverse from every entry but the one starting w shiny bag
	// when done, if shiny is in the list => +1

	var containsShiny []int

	for index, bag := range bags {
		if bag.color != "shiny gold" {
			metShiny := traverse(bags, bag, index)
			if metShiny {
				containsShiny = append(containsShiny, index)
			}
		}
	}
	return len(containsShiny)

	// traverse(bags, bags[0])
}

func traverse(bags []Bag, current Bag, index int) bool {
	var metShiny = false
	if (len(current.contains) == 0 || metShiny) {
		// now we're done.
		fmt.Printf("reached shiny gold or no more bags\n")
		return true
	} else {
		for _, sub := range current.contains {
			if sub.color == "shiny gold" {
				metShiny = true
				fmt.Printf("met shiny, breaking\n")
				break
			} else {
				// traverse further
				fmt.Printf("here [%v]\n", sub)
				// UGLY need to find the Bag holding sub.color
				for _, next := range bags {
					if sub.color == next.color {
						traverse(bags, next, index)
					}
				}
			}
		}
	}
	return metShiny
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
	fmt.Printf("PART 1 %v\n", findShinyBags(rules))
}
