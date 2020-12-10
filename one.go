package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var n = flag.Int("n", 0, "How many add up to 2020?")
	flag.Parse()

	input := []int{}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		i, _ := strconv.Atoi(sc.Text())
		input = append(input, i)
	}

	if *n == 3 {
		result := runThree(input)
		fmt.Println(result)
	} else {
		result := runTwo(input)
		fmt.Println(result)
	}
}

func runTwo(input []int) int {
	for _, x := range input {
		for _, y := range input {
			if x+y == 2020 {
				return x * y
			}
		}
	}

	log.Fatal("Didn't work")
	return 0
}

func runThree(input []int) int {
	for _, x := range input {
		for _, y := range input {
			for _, z := range input {
				if x+y+z == 2020 {
					return x * y * z
				}
			}
		}
	}

	log.Fatal("Didn't work")
	return 0
}
