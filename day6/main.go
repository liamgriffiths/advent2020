package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	var part2 = flag.Bool("2", false, "")
	flag.Parse()

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
			input[n] = append(input[n], line)
		}
	}

	sum := 0
	for _, group := range input {
		if *part2 {
			sum = sum + allSaidYes(group)
		} else {
			sum = sum + atLeastOneYesEach(group)
		}
	}

	fmt.Println(sum)
}

// part 1
func atLeastOneYesEach(group []string) int {
	ys := make(map[string]bool)
	for _, answers := range group {
		for _, answer := range answers {
			ys[string(answer)] = true
		}
	}

	return len(ys)
}

// part 2
func allSaidYes(group []string) int {
	ys := make(map[string][]bool)
	for _, answers := range group {
		for _, answer := range answers {
			ys[string(answer)] = append(ys[string(answer)], true)
		}
	}

	count := 0
	for _, v := range ys {
		if len(v) == len(group) {
			count = count + 1
		}
	}

	return count
}
