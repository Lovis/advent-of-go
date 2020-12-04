package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// type Passport struct {
// 	byr string // (Birth Year)
// 	iyr string // (Issue Year)
// 	eyr string // (Expiration Year)
// 	hgt string // (Height)
// 	hcl string // (Hair Color)
// 	ecl string // (Eye Color)
// 	pid string // (Passport ID)
// 	cid string // (Country ID) optional
// }

// break on empty line

func main() {
	fmt.Printf("Day 4\n")

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Printf("error opening file")
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var passports [][]string

	var passport []string
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			passports = append(passports, passport)
			passport = []string{}
		} else {
			attributes := strings.Split(s, " ")
			passport = append(passport, attributes...)
		}
	}
	fmt.Printf("reading passports: %v\n", passports)
}
