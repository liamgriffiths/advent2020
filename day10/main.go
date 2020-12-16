package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	joltage := []int{}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		j, _ := strconv.Atoi(sc.Text())
		joltage = append(joltage, j)
	}

	part2(joltage)
}

func part1(js []int) {
	// add the seat plug
	js = append(js, 0)
	sort.Ints(js)
	// add the device
	js = append(js, js[len(js)-1]+3)

	out := make(map[int]int)

	for i, n := range js {
		if i+1 < len(js) {
			diff := js[i+1] - n
			out[diff] = out[diff] + 1
		}
	}

	fmt.Println(out[1] * out[3])
}

func part2(js []int) {
	// todo
}
