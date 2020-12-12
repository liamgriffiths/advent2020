package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := [][]string{}
	input = append(input, []string{})
	n := 0

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			n = n + 1
			input = append(input, []string{})
		} else {
			fs := strings.Fields(line)
			if len(fs) > 0 {
				input[n] = append(input[n], fs...)
			}
		}
	}

	passports := []passport{}
	for _, in := range input {
		passports = append(passports, parse(in))
	}

	validCount := 0
	for _, p := range passports {
		if p.isValid() {
			validCount = validCount + 1
		}
	}

	fmt.Println(validCount)
}

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p *passport) isValid() bool {
	if p.byr != "" &&
		p.iyr != "" &&
		p.eyr != "" &&
		p.hgt != "" &&
		p.hcl != "" &&
		p.ecl != "" &&
		p.pid != "" {

		// byr (Birth Year) - four digits; at least 1920 and at most 2002.
		byr, _ := strconv.Atoi(p.byr)
		if byr < 1920 || byr > 2002 {
			return false
		}

		// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
		iyr, _ := strconv.Atoi(p.iyr)
		if iyr < 2010 || iyr > 2020 {
			return false
		}

		// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
		eyr, _ := strconv.Atoi(p.eyr)
		if eyr < 2020 || byr > 2030 {
			return false
		}

		// hgt (Height) - a number followed by either cm or in:
		// If cm, the number must be at least 150 and at most 193.
		// If in, the number must be at least 59 and at most 76.
		if len(p.hgt) < 2 {
			return false
		}
		hgtu := p.hgt[len(p.hgt)-2:]
		hgtn, _ := strconv.Atoi(p.hgt[:len(p.hgt)-2])
		if hgtu == "cm" {
			if hgtn < 150 || hgtn > 193 {
				return false
			}
		} else if hgtu == "in" {
			if hgtn < 59 || hgtn > 76 {
				return false
			}
		} else {
			return false
		}

		// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
		if string(p.hcl[0]) == "#" && len(p.hcl[1:]) == 6 {
			isOk := regexp.MustCompile(`^[a-f0-9]+$`).MatchString
			if !isOk(p.hcl[1:]) {

				return false
			}
		} else {

			return false
		}

		// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
		if !(p.ecl == "amb" || p.ecl == "blu" || p.ecl == "brn" || p.ecl == "gry" || p.ecl == "grn" || p.ecl == "hzl" || p.ecl == "oth") {
			return false
		}

		// pid (Passport ID) - a nine-digit number, including leading zeroes.
		if len(p.pid) != 9 {
			return false
		}
		isOk := regexp.MustCompile(`^[0-9]+$`).MatchString
		if !isOk(p.pid) {
			return false
		}

		// cid (Country ID) - ignored, missing or not.
		return true
	}

	return false
}

func parse(input []string) passport {
	newPassport := passport{}

	for _, token := range input {
		s := strings.Split(token, ":")
		name := s[0]
		value := s[1]

		if name == "byr" {
			if newPassport.byr == "" {
				newPassport.byr = value
			}
		}

		if name == "iyr" {
			if newPassport.iyr == "" {
				newPassport.iyr = value
			}
		}

		if name == "eyr" {
			if newPassport.eyr == "" {
				newPassport.eyr = value
			}
		}

		if name == "hgt" {
			if newPassport.hgt == "" {
				newPassport.hgt = value
			}
		}

		if name == "hcl" {
			if newPassport.hcl == "" {
				newPassport.hcl = value
			}
		}

		if name == "ecl" {
			if newPassport.ecl == "" {
				newPassport.ecl = value
			}
		}

		if name == "pid" {
			if newPassport.pid == "" {
				newPassport.pid = value
			}
		}

		if name == "cid" {
			if newPassport.cid == "" {
				newPassport.cid = value
			}
		}
	}

	return newPassport
}
