package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var part2 = flag.Bool("2", false, "Part 2?")
	flag.Parse()

	entries := []Entry{}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		x := sc.Text()
		t := parse(x)
		entries = append(entries, *t)
	}

	validCount := 0

	for _, e := range entries {
		if *part2 {
			if e.isValidByPosition() {
				validCount = validCount + 1
			}
		} else {
			if e.isValidByRange() {
				validCount = validCount + 1
			}
		}
	}

	fmt.Println(validCount)
}

type Entry struct {
	Min      int
	Max      int
	Char     string
	Password string
}

func (e *Entry) isValidByRange() bool {
	count := 0
	for _, c := range e.Password {
		if string(c) == e.Char {
			count = count + 1
		}
	}
	return count >= e.Min && count <= e.Max
}

func (e *Entry) isValidByPosition() bool {
	x := e.Password[e.Min-1]
	y := e.Password[e.Max-1]

	if string(x) == e.Char || string(y) == e.Char {
		return string(x) != string(y)
	} else {
		return false
	}
}

func parse(line string) *Entry {
	re := regexp.MustCompile(`(?P<min>\d+)-(?P<max>\d+) (?P<char>.): (?P<password>.*)`)
	names := re.SubexpNames()
	y := re.FindAllStringSubmatch(line, -1)[0]

	kv := map[string]string{}
	for i, n := range y {
		kv[names[i]] = n
	}

	min, _ := strconv.Atoi(kv["min"])
	max, _ := strconv.Atoi(kv["max"])

	return &Entry{
		Min:      min,
		Max:      max,
		Char:     kv["char"],
		Password: kv["password"],
	}
}
