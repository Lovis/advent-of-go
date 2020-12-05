package main

import (
	"bufio"
	"fmt"
	"os"
)

// position
// 128 rows
// 8 cols
type Seat struct {
	row int
	col int
	id  int
}

func main() {
	fmt.Printf("Day 5\n")

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Printf("error opening file")
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	// var seats []Seat

	// var highest = 0
	var ROWS [128]int
	for i := range ROWS {
		ROWS[i] = i
	}

	var COLS [8]int
	for i := range COLS {
		COLS[i] = i
	}

	var seats []Seat
	var highest = 0

	for scanner.Scan() {
		input := scanner.Text()

		row := input[0:8]
		col := input[7:10]

		var rowUpper = len(ROWS)
		var rowLower = 0
		// for every char in row, move the cursor
		for _, char := range row {
			if string(char) == "F" {
				rowUpper -= (rowUpper - rowLower) / 2
			} else {
				rowLower += (rowUpper - rowLower) / 2
			}
		}

		var colUpper = len(COLS)
		var colLower = 0
		// for every char in row, move the cursor
		for _, char := range col {
			if string(char) == "L" {
				colUpper -= (colUpper - colLower) / 2
			} else {
				colLower += (colUpper - colLower) / 2
			}
		}
		var seat Seat
		seat.row = ROWS[rowLower:rowUpper][0]
		seat.col = COLS[colLower:colUpper][0]
		seat.id = seat.row*8 + seat.col

		if seat.id > highest {
			highest = seat.id
		}
		seats = append(seats, seat)
	}

	fmt.Printf("seats: %v, highest: %d\n", seats, highest)
}
