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

func traversing(bags []Bag) int {
	var holdingShiny []int

	for initialIndex, initialBag := range bags {
		if (initialBag.color == "shiny bag") {
			continue
		} else {
			found := recursive(bags, initialBag, holdingShiny, initialIndex)
			if found {
				fmt.Printf("returned true, %v, %v\n", initialIndex, holdingShiny)
				holdingShiny = append(holdingShiny, initialIndex)
			}
		}
	}
	// holdingShiny := recursive(bags, bags[0], []int{} , 0)
	return len(holdingShiny)
}

func recursive(bags []Bag, current Bag, indicies []int, initialIndex int) bool {
	fmt.Printf("incoming: %v\n", indicies)
	found := false
	for _, index := range indicies {
		if index == initialIndex {
			found = true
		}
	}
	// fmt.Printf("new call: bag: %v, found: %t\n", current, found)
	if current.color == "shiny gold" {
		fmt.Printf("++ exiting!\n")
		found = true
		indicies = append(indicies, initialIndex)
		fmt.Printf("indicies %v\n", indicies)
		// return indicies
	}
	if (!found && len(current.contains) > 0) {
		// go deeper
		for _, child := range current.contains {
			if found || child.color == "shiny gold" {
				found = true
				indicies = append(indicies, initialIndex)
				fmt.Printf("!! met shiny, breaking\n")
				fmt.Printf("indicies %v\n", indicies)
				// return indicies
				break
			}
			// UGLY need to find the Bag holding sub.color
			for _, next := range bags {
				if child.color == next.color {
					fmt.Printf("-- child [%v]\n", next.color)
					recursive(bags, next, indicies, initialIndex)
				}
			}
		}
	}
	return found
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
	// for _, rule := range rules {
	// 	fmt.Printf("%+v\n", rule)
	// }
	// fmt.Printf("PART 1 %v\n", findShinyBags(rules))
	fmt.Printf("PART 1: count %v\n", traversing(rules))
}
