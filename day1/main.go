package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var part2 = flag.Bool("2", false, "")
	flag.Parse()

	input := []int{}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		i, _ := strconv.Atoi(sc.Text())
		input = append(input, i)
	}

	if *part2 {
		fmt.Println(runPart2(input))
	} else {
		fmt.Println(runPart1(input))
	}
}

func runPart1(input []int) int {
	for _, x := range input {
		for _, y := range input {
			if x+y == 2020 {
				return x * y
			}
		}
	}

	return 0
}

func runPart2(input []int) int {
	for _, x := range input {
		for _, y := range input {
			for _, z := range input {
				if x+y+z == 2020 {
					return x * y * z
				}
			}
		}
	}

	return 0
}
