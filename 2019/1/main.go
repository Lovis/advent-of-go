package main

import (
	"fmt"
	"math"
)

var data = [100]int{
118868,
88841,
133680,
148066,
70887,
93213,
124243,
92767,
71322,
86793,
53650,
102703,
146958,
53031,
148282,
124989,
74375,
122044,
122693,
74204,
74869,
81803,
124436,
68495,
74865,
70765,
81537,
61376,
145342,
137159,
115230,
119293,
147126,
130191,
131330,
122891,
135407,
116334,
130325,
138521,
71955,
53806,
122260,
102573,
70032,
75981,
111555,
135654,
50805,
122186,
138172,
96422,
124781,
55894,
54337,
149926,
63809,
146163,
55131,
55796,
92771,
80288,
111619,
134602,
82245,
72505,
117209,
92383,
149101,
135399,
112166,
134000,
88771,
63963,
103731,
74915,
146347,
125390,
126249,
131534,
142038,
55327,
58784,
85003,
65909,
89879,
128715,
138559,
146209,
145040,
116032,
130046,
131664,
125899,
141918,
88426,
50488,
67943,
79677,
94858}

var testData = [2]int{14, 1969}

func calc (x int) int {
	divided := math.Floor(float64(x) / 3)
	return int(divided) - 2
}

func main() {
	fmt.Printf("HELLO here worjd %v\n", testData[0])
	// part2()
	part2()
	
}

func part1 () {
	var total = 0

	for i := 0; i < 100; i++ {
		total += calc(data[i])
	}
	fmt.Printf("Total sum of fuel: %d \n", total)
		// 3457681
}

func negativeFuel (x int) int {
	var fuel = calc(x)
	fmt.Printf("fuel: %d \n", fuel)
	if (calc(fuel) <= 0) {
		fmt.Printf("less than 0!\n")
		return fuel
	}
	return fuel + negativeFuel(fuel)
}

func part2 () {
	var total = 0

	for i := 0; i < 100; i++ {
		total += negativeFuel(data[i])
	}
	fmt.Printf("Total sum of fuel: %d \n", total)
	// 5183653
}

