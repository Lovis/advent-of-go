package main

import (
	"fmt"
	"strconv"
	"strings"
)


type Limit struct {
	lower, upper int
}

func hasDouble(number int) (bool) {
	converted := strconv.Itoa(number)
	hasDouble := false
	for i := 0; i < len(converted); i++ {
		// rule 1: increasing
		// fmt.Println(i, string(converted[i]))
		if (i > 0 && string(converted[i]) < string(converted[i-1])) {
			hasDouble = false
			return hasDouble
		}

		// rule 2: mnultiple occurrencies
		sep := string(converted[i])
		length := len(strings.Split(converted, sep))
		if (length > 2) {
			hasDouble = true
		}
		
	}
	return hasDouble
}

func main() {
	fmt.Println("Puzzle 4")
	fmt.Println("has double? ", hasDouble(123456))
	fmt.Println("has double? ", hasDouble(123356))
	fmt.Println("has double? ", hasDouble(153356))

	fmt.Println("has double? ", hasDouble(111111))
	fmt.Println("has double? ", hasDouble(223450))
	fmt.Println("has double? ", hasDouble(123789))
	// config := Limit{136760, 595730}

}