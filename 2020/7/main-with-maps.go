package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Bag a littl struct for you
type Bag struct {
	color string
	childs map[*Bag]int
}

func findShinyBag(bags []Bag) int {
	var myBag = "shiny gold"

	for _, bag := range bags {
		if (bag.color != myBag) {
			search(bags, bag, myBag)
		}
	}
	return 0
}

func search(bags []Bag, current Bag, myBag string) bool {
	if current.color == myBag {
		return true
	}

	// if (len(current.childs) == 0) {
	// 	// dead end
	// 	return false
	// }

	// "red": 4
	// for color, count := range current.childs {
	// 	search(bags, )

	// 	for _, bag := range bags {
	// 		if bag.color == color {
	// 			search(bags, bag, myBag)
	// 		}
	// 	}
	// }
	return false
}

func parseInput(scanner *bufio.Scanner) map[string]Bag {
	bags := make(map[string]Bag)
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
			// fmt.Printf("current col: %s\n", color)

			// find or create create a bag pointer for the child
			var child Bag
			for _, existing := range bags {
				if existing.color == color {
					child = existing
				}
			}

			child.color = color
			// fmt.Printf("color %s count %d\n", child.color, count)
			children[&child] = count

			if _, ok := bags[child.color]; ok {
			} else {
				bags[child.color] = child
			}
		}


		bag.childs = children
		if _, ok := bags[bag.color]; ok {
		} else {
			bags[bag.color] = bag
		}
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

	var bags = parseInput(scanner)
	fmt.Printf("PART 1, bags:\n")
	for _, bag := range bags {
		fmt.Printf("bag: %v\n", bag.color)
		for key, value := range bag.childs {
			fmt.Printf("-- child: %v %v\n", key, value)
		}
	}
}
