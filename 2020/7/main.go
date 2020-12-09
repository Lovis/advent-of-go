package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// 	"strconv"
// )

// type Bag struct {
// 	Color string
// 	Childs map[string]*Bag
// }

// func NewBag(Color string) *Bag {
// 	return &Bag {
// 		Color: Color,
// 		Childs: map[string]*Bag{},
// 	}
// }

// type Graph struct {
// 	Childs map[string]*Bag
// }

// func NewGraph() *Graph {
// 	return &Graph{
// 		Childs: map[string]*Bag{},
// 	}
// }

// func (g *Graph) AddBag(color string) {
// 	v := NewBag(color)
// 	g.Childs[color] = v
// }

// func (g *Graph) AddEdge(k1, k2 string, count int) {
// 	v1 := g.Childs[k1]
// 	v2 := g.Childs[k2]

// 	if _, ok := v1.Childs[v2.Color]; ok {
// 		return
// 	}

// 	v1.Childs[v2.Color] = v2
// 	g.Childs[v1.Color] = v1
// 	g.Childs[v2.Color] = v2

// }
// // type Children struct {
// // 	number int
// // 	color string
// // }

// // type Bag struct {
// // 	color string
// // 	children map[Children]int
// // }

// // func parseRules(scanner *bufio.Scanner) []Bag {
// // 	var bags []Bag
// // 	for scanner.Scan() {
// // 		var bag Bag
// // 		s := scanner.Text()
// // 		if s == "" {
// // 			continue
// // 		}
// // 		parts := strings.Split(s, "contain ")
// // 		bag.color = strings.Trim(parts[0], " bags")
// // 		contains := strings.Split(parts[1], ", ")

// // 		var children map[]
// // 		// content: "1 bright white bag"
// // 		// "no other bag"
// // 		for _, content := range contains {
// // 			var c Children
// // 			if content == "no other bags." {
// // 				continue
// // 			}
// // 			splits := strings.Split(content, " ")
// // 			c.number, _ = strconv.Atoi(splits[0])
// // 			c.color = strings.Join(splits[1:len(splits)-1], " ")
// // 			contents = append(contents, c)
// // 		}
// // 		bag.children = contents
// // 		bags = append(bags, bag)
// // 	}

// // 	return bags
// // }

// // func traversing(bags []Bag) int {
// // 	var holdingShiny []int

// // 	for initialIndex, initialBag := range bags {
// // 		if (initialBag.color == "shiny bag") {
// // 			continue
// // 		} else {
// // 			found := recursive(bags, initialBag, holdingShiny, initialIndex)
// // 			if found {
// // 				fmt.Printf("returned true, %v, %v\n", initialIndex, holdingShiny)
// // 				holdingShiny = append(holdingShiny, initialIndex)
// // 			}
// // 		}
// // 	}
// // 	// holdingShiny := recursive(bags, bags[0], []int{} , 0)
// // 	return len(holdingShiny)
// // }

// // func recursive(bags []Bag, current Bag, indicies []int, initialIndex int) bool {
// // 	fmt.Printf("incoming: %v\n", indicies)
// // 	found := false
// // 	for _, index := range indicies {
// // 		if index == initialIndex {
// // 			found = true
// // 		}
// // 	}
// // 	// fmt.Printf("new call: bag: %v, found: %t\n", current, found)
// // 	if current.color == "shiny gold" {
// // 		fmt.Printf("++ exiting!\n")
// // 		found = true
// // 		indicies = append(indicies, initialIndex)
// // 		fmt.Printf("indicies %v\n", indicies)
// // 		// return indicies
// // 	}
// // 	if (!found && len(current.contains) > 0) {
// // 		// go deeper
// // 		for _, child := range current.contains {
// // 			if found || child.color == "shiny gold" {
// // 				found = true
// // 				indicies = append(indicies, initialIndex)
// // 				fmt.Printf("!! met shiny, breaking\n")
// // 				fmt.Printf("indicies %v\n", indicies)
// // 				// return indicies
// // 				break
// // 			}
// // 			// UGLY need to find the Bag holding sub.color
// // 			for _, next := range bags {
// // 				if child.color == next.color {
// // 					fmt.Printf("-- child [%v]\n", next.color)
// // 					recursive(bags, next, indicies, initialIndex)
// // 				}
// // 			}
// // 		}
// // 	}
// // 	return found
// // }

// type Relation struct {
// 	parent string
// 	child string
// 	count int
// }
// func main1() {
// 	fmt.Printf("Day 7\n")

// 	file, err := os.Open("tiny.txt")

// 	if err != nil {
// 		fmt.Printf("error opening file\n")
// 		return
// 	}

// 	scanner := bufio.NewScanner(file)

// 	scanner.Split(bufio.ScanLines)

// 	var bags []string
// 	var relations []Relation
// 	for scanner.Scan() {
// 		var bag string
// 		s := scanner.Text()
// 		if s == "" {
// 			continue
// 		}
// 		parts := strings.Split(s, "contain ")
// 		bag = strings.Trim(parts[0], " bags")
// 		contains := strings.Split(parts[1], ", ")


// 		// content: "1 bright white bag"
// 		// "no other bag"
// 		for _, content := range contains {
// 			var c Relation
// 			if content == "no other bags." {
// 				continue
// 			}
// 			splits := strings.Split(content, " ")
// 			c.parent = bag
// 			c.count, _ = strconv.Atoi(splits[0])
// 			c.child = strings.Join(splits[1:len(splits)-1], " ")
// 			relations = append(relations, c)
// 		}
// 		bags = append(bags, bag)

// 		// fill the graph

// 		g := NewGraph()

// 		for _, color := range bags {
// 			g.AddBag(color)
// 		}

// 		for _, relation := range relations {
// 			g.AddEdge(relation.parent, relation.child, relation.count)
// 		}
// 	}
// }
