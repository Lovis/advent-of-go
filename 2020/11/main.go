package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func getAdjacentOccupiedSeats(cy int, cx int, seatmap map[int][]string) int {
	occupiedSeats := 0

	for y := cy - 1; y <= cy+1; y++ {
		for x := cx - 1; x <= cx+1; x++ {
			if y == cy && x == cx {
				continue
			}
			if row, ok := seatmap[y]; ok {
				if x >= 0 && x < len(row) {
					neighbor := row[x]
					if (neighbor) == "#" {
						occupiedSeats++
					}
				}
			}
		}
	}

	return occupiedSeats
}

func simulate(seatmap map[int][]string) map[int][]string {
	var simulated = make(map[int][]string)
	var changed = false
	// cred: https://gist.github.com/cjgiridhar/f95a5f6b1890d9a4b9ac0a8343806bdd
	jsonStr, _ := json.Marshal(seatmap)
	json.Unmarshal(jsonStr, &simulated)

	// fmt.Printf("\ninput: \n")
	for i := 0; i < len(seatmap); i++ {
		// fmt.Printf("%d %v\n", i, seatmap[i])
	}

	for i, row := range seatmap {
		for k, seat := range row {
			// check neighbors

			occupiedSeats := getAdjacentOccupiedSeats(i, k, seatmap)
			// fmt.Printf("checking %d %d: neighbors: %d\n", i, k, occupiedSeats)
			if seat == "L" && occupiedSeats == 0 {
				changed = true
				simulated[i][k] = "#"
				// fmt.Printf("changing to # %d %d. input: %s changed: %s\n", i, k, seat, simulated[i][k])
			} else if seat == "#" && occupiedSeats >= 4 {
				changed = true
				simulated[i][k] = "L"
				// fmt.Printf("changing to L %d %d input: %s changed: %s \n", i, k, seat, simulated[i][k])
			}
		}
	}
	if changed {
		return simulate(simulated)
	}
	return simulated
}
func parseInput() map[int][]string {
	file, err := os.Open("input.txt")
	var seatmap = make(map[int][]string)

	if err != nil {
		fmt.Printf("error opening file\n")
		return seatmap
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	index := 0
	for scanner.Scan() {
		var row []string
		s := scanner.Text()
		if s == "" {
			continue
		}
		row = strings.Split(s, "")
		seatmap[index] = row
		index++
	}
	return seatmap
}

func main() {

	fmt.Printf("Day 11\n")

	seatmap := parseInput()

	fmt.Printf("part 1\n")
	// fmt.Printf("input: \n")
	// for i := 0; i < len(seatmap); i++ {
	// 	fmt.Printf("%v\n", seatmap[i])
	// }

	simulated := simulate(seatmap)
	fmt.Printf("\n 1. simulate: \n")
	count := 0
	for i := 0; i < len(simulated); i++ {
		fmt.Printf("%d %v\n", i, simulated[i])
		for _, seat := range simulated[i] {
			if seat == "#" {
				count++
			}
		}
	}
	fmt.Printf("occupied: %d\n", count)
}
