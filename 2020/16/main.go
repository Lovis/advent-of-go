package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var notes = make(map[string][]int)
var yourticket []int
var tickets [][]int

var valid = make(map[string]map[int]int)
var solution = make(map[int]string)
var options = make(map[string][]int)

func applyPossibles() {
	// row: [0 1 2], seat: [2]
	for name, posList := range options {
		var relevant []int
		// then, check array
		for _, pos := range posList {
			if solution[pos] == "" {
				relevant = append(relevant, pos)
			}
		}
		if len(relevant) == 1 {
			solution[relevant[0]] = name
		}
	}
	if len(solution) == len(notes) {
		return
	}
	applyPossibles()
}
func setSolution() {
	for name, valids := range valid {
		var cols []int
		for key, value := range valids {
			if value == len(tickets) {
				cols = append(cols, key)
			}
		}
		options[name] = cols
	}
	applyPossibles()
}

func part2() int {
	for _, ticket := range tickets {
		for index, number := range ticket {
			for name, rule := range notes {
				nameKey := valid[name]
				if nameKey == nil {
					nameKey = make(map[int]int)
				}

				if (number >= rule[0] && number <= rule[1]) || (number >= rule[2] && number <= rule[3]) {
					nameIndex := nameKey[index]
					nameIndex++
					nameKey[index] = nameIndex
				}
				valid[name] = nameKey
			}
		}
	}

	setSolution()
	sum := 1
	for index, name := range solution {
		if strings.Contains(name, "departure") {
			sum *= yourticket[index]
		}
	}

	return sum
}

func part1() int {
	sum := 0
	var invalidIds []int
	for index, ticket := range tickets {
		var valid = make(map[string][]int)

		var usedNumbers = make(map[int]bool)
		for _, number := range ticket {
			for name, rule := range notes {
				values := valid[name]

				if (number >= rule[0] && number <= rule[1]) || (number >= rule[2] && number <= rule[3]) {
					values = append(values, number)
					usedNumbers[number] = true
				}
				valid[name] = values
			}
		}
		if len(usedNumbers) < len(ticket) {

			for _, number := range ticket {
				if _, ok := usedNumbers[number]; !ok {
					invalidIds = append([]int{index}, invalidIds...)
					sum += number
				}
			}
		}
	}

	// delete invalid tickets
	for _, inv := range invalidIds {
		tickets = append(tickets[:inv], tickets[inv+1:]...)
	}
	return sum
}

func parseInput() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening file\n")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var section = 0
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}
		if s == "your ticket:" || s == "nearby tickets:" {
			section = section + 1
			continue
		}
		if section == 0 {
			key := strings.Split(s, ":")[0]
			var ints []int
			values := strings.Trim(strings.Split(s, ":")[1], " ")
			for _, numbers := range strings.Split(values, "or") {
				vars := strings.Split(strings.Trim(numbers, " "), "-")
				for _, v := range vars {
					vi, _ := strconv.Atoi(strings.Trim(v, ""))
					ints = append(ints, vi)
				}
			}
			notes[key] = ints
		} else if section == 1 {
			numbers := strings.Split(s, ",")
			for _, number := range numbers {
				value, _ := strconv.Atoi(number)
				yourticket = append(yourticket, value)
			}
		} else if section == 2 {
			var ticket []int
			numbers := strings.Split(s, ",")
			for _, number := range numbers {
				value, _ := strconv.Atoi(number)
				ticket = append(ticket, value)
			}
			tickets = append(tickets, ticket)
		}
	}
}

func main() {
	fmt.Printf("hello, day 16\n")

	parseInput()
	res := part1()
	fmt.Printf("Part 1: %d\n", res)
	res2 := part2()
	fmt.Printf("Part 2: %d\n", res2)
}
