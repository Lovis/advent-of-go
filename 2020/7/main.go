package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var bags = make(map[string]Bag)
// Bag a littl struct for you
type Bag struct {
	color string
	childs map[*Bag]int
}

func findShinyBag() int {
	var myBag = "shiny gold"
	var counter = 0
	for _, bag := range bags {
		if (bag.color != myBag) {
			found := search(bag, myBag)
			if found == true {
				counter++
			}
		}
	}
	// fmt.Printf("res %t\n", search(bags["dark brown"], myBag))
	return counter
}

func search(current Bag, myBag string) bool {
	// fmt.Printf("\nbag %v\n", current)
	if current.color == myBag {
		// fmt.Printf("found gold %v\n", current.color)
		return true
	}

	// this makes one branch only
	for childBag := range current.childs {
		// fmt.Printf("running childbag %v\n", childBag)
		found := search(bags[childBag.color], myBag)
		if found {
			return true
		}

	}
	return false
}

func parseInput(scanner *bufio.Scanner) map[string]Bag {
	for scanner.Scan() {
		var bag Bag

		s := scanner.Text()
		if s == "" {
			continue
		}
		parts := strings.Split(s, "contain ")
		bag.color = strings.Split(parts[0], " bags")[0]

		childs := strings.Split(parts[1], ", ")

		children := make(map[*Bag]int)

		for _, child := range childs {
			if child == "no other bags." {
				continue
			}
			splits := strings.Split(child, " ")
			count, _ := strconv.Atoi(splits[0])
			color := strings.Join(splits[1:len(splits)-1], " ")

			var child Bag
			for _, existing := range bags {
				if existing.color == color {
					child = existing
				}
			}

			child.color = color
			children[&child] = count

			if _, ok := bags[child.color]; ok {
			} else {
				bags[child.color] = child
			}
		}


		bag.childs = children
		bags[bag.color] = bag
	}

	return bags
}

func main() {
	fmt.Printf("Day 7\n")

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Printf("error opening file\n")
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	parseInput(scanner)
	// fmt.Printf("PART 1, bags:\n")
	// for _, bag := range bags {
	// 	fmt.Printf("bag: %v\n", bag.color)
	// 	for key, value := range bag.childs {
	// 		fmt.Printf("-- child: %v %v\n", key, value)
	// 	}
	// }

	fmt.Printf("search: %d\n", findShinyBag())
}
