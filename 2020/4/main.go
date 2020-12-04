package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
var eyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func ecl(color string) bool {
	var valid = false
	for _, eyeColor := range eyeColors {
		if color == eyeColor {
			valid = true
			break
		}
	}
	return valid
}

func verifiedPassport(passport []string) bool {
	var cid = false
	var presentParameters = false
	var validParameters = true

	for _, attribute := range passport {
		keyvalue := strings.Split(attribute, ":")
		key, value := keyvalue[0], keyvalue[1]

		switch key {
		case "byr":
			year, _ := strconv.Atoi(value)
			// fmt.Printf("year %d %t %t\n", year, 1920 <= year, 2002 >= year)
			if !(1920 <= year && 2002 >= year) {
				validParameters = false
				break
			}

		case "iyr":
			year, _ := strconv.Atoi(value)
			if !(2010 <= year && 2020 >= year) {
				validParameters = false
				break
			}

		case "eyr":
			year, _ := strconv.Atoi(value)
			if !(2020 <= year && 2030 >= year) {
				validParameters = false
				break
			}

		case "hgt":
			if strings.Contains(value, "in") {
				s := strings.Split(value, "in")
				var height, _ = strconv.Atoi(s[0])
				if !(59 <= height && 76 >= height) {
					validParameters = false
					break
				}
			} else if strings.Contains(value, "cm") {
				s := strings.Split(value, "cm")
				var height, _ = strconv.Atoi(s[0])
				if !(150 <= height && 193 >= height) {
					validParameters = false
					break
				}
			}

		case "hcl":
			matched, _ := regexp.MatchString(`^#[a-f0-9]{6}$`, value)
			if !matched {
				validParameters = false
				break
			}

		case "ecl":
			validEyeColor := ecl(value)
			if !validEyeColor {
				validParameters = false
				break
			}

		case "pid":
			matched, _ := regexp.MatchString(`^[0-9]{9}`, value)
			if !matched {
				validParameters = false
				break
			}
		}

		if key == "cid" {
			cid = true
			continue
		}
	}

	presentParameters = len(passport) == 8 || (len(passport) == 7 && !cid)
	return validParameters && presentParameters
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
			if verifiedPassport(passport) {
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
