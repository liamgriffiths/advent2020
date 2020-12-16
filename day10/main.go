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

	// add the seat plug
	joltage = append(joltage, 0)
	sort.Ints(joltage)
	// add the device
	joltage = append(joltage, joltage[len(joltage)-1]+3)

	joltages(joltage)
	// fmt.Println(joltage)
}

func joltages(js []int) {
	out := make(map[int]int)

	for i, n := range js {
		if i+1 < len(js) {
			diff := js[i+1] - n
			out[diff] = out[diff] + 1
		}
	}

	fmt.Println(out[1] * out[3])
}
