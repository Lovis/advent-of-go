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

func validPassport(passport []string) bool {
	var cid = false
	for _, attribute := range passport {
		if strings.Split(attribute, ":")[0] == "cid" {
			cid = true
			continue
		}
	}
	// fmt.Printf("valid? %d %t\n", len(passport), cid)
	return len(passport) == 8 || (len(passport) == 7 && !cid)
}

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
			// end of passport.
			// verify length.
			// if length is less than 1, look
			if validPassport(passport) {
				passports = append(passports, passport)
			}
			passport = []string{}
		} else {
			attributes := strings.Split(s, " ")
			passport = append(passport, attributes...)
		}
	}

	fmt.Printf("valid passports: %v\n", len(passports))
}
